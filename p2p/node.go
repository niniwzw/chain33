package p2p

import (
	"fmt"
	"math/rand"

	"os"
	"sync"
	"time"

	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/types"
)

// 启动Node节点
//1.启动监听GRPC Server
//2.初始化AddrBook,并发起对相应地址的连接
//3.如果配置了种子节点，则连接种子节点
//4.启动监控远程节点
func (n *Node) Start() {
	n.DetectNodeAddr()
	n.l = NewDefaultListener(Protocol, n)
	go n.monitor()
	go n.doNat()

	return
}

func (n *Node) Close() {

	close(n.loopDone)
	log.Debug("stop", "versionDone", "closed")
	if n.l != nil {
		n.l.Close()
	}

	log.Debug("stop", "listen", "closed")
	n.addrBook.Close()
	log.Debug("stop", "addrBook", "closed")
	n.RemoveAll()
	log.Debug("stop", "PeerRemoeAll", "closed")

}

type Node struct {
	omtx     sync.Mutex
	addrBook *AddrBook // known peers
	nodeInfo *NodeInfo
	outBound map[string]*peer
	l        Listener
	loopDone chan struct{}
}

func (n *Node) SetQueue(q *queue.Queue) {
	n.nodeInfo.q = q
	n.nodeInfo.qclient = q.NewClient()
}

func NewNode(cfg *types.P2P) (*Node, error) {
	os.MkdirAll(cfg.GetDbPath(), 0755)
	node := &Node{
		outBound: make(map[string]*peer),
		addrBook: NewAddrBook(cfg.GetDbPath() + "/addrbook.json"),
		loopDone: make(chan struct{}, 1),
	}

	node.nodeInfo = NewNodeInfo(cfg)

	return node, nil
}
func (n *Node) FlushNodePort(localport, export uint16) {

	if exaddr, err := NewNetAddressString(fmt.Sprintf("%v:%v", n.nodeInfo.GetExternalAddr().IP.String(), export)); err == nil {
		n.nodeInfo.SetExternalAddr(exaddr)
	}
	if listenAddr, err := NewNetAddressString(fmt.Sprintf("%v:%v", LocalAddr, localport)); err == nil {
		n.nodeInfo.SetListenAddr(listenAddr)
	}

}

func (n *Node) NatOk() bool {
	n.nodeInfo.natNoticeChain <- struct{}{}
	ok := <-n.nodeInfo.natResultChain
	return ok
}

func (n *Node) doNat() {

	if OutSide == false { //如果能作为服务方，则Nat,进行端口映射，否则，不启动Nat

		if !n.NatOk() {
			SERVICE -= NODE_NETWORK //nat 失败，不对外提供服务
			log.Info("doNat", "NatFaild", "No Support Service")
		} else {
			log.Info("doNat", "NatOk", "Support Service")
		}

	}
	n.addrBook.AddOurAddress(n.nodeInfo.GetExternalAddr())
	n.addrBook.AddOurAddress(n.nodeInfo.GetListenAddr())
	if selefNet, err := NewNetAddressString(fmt.Sprintf("127.0.0.1:%v", n.nodeInfo.GetListenAddr().Port)); err == nil {
		n.addrBook.AddOurAddress(selefNet)
	}
	close(n.nodeInfo.natDone)
	return
}

func (n *Node) DialPeers(addrbucket map[string]bool) error {
	if len(addrbucket) == 0 {
		return nil
	}
	var addrs []string
	for addr, _ := range addrbucket {
		addrs = append(addrs, addr)
	}
	netAddrs, err := NewNetAddressStrings(addrs)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	for _, netAddr := range netAddrs {
		//log.Info("DialPeers", "netAddr", netAddr)
		if n.addrBook.ISOurAddress(netAddr) == true {
			continue
		}

		//不对已经连接上的地址重新发起连接
		if n.Has(netAddr.String()) {
			log.Debug("DialPeers", "find hash", netAddr.String())
			continue
		}

		if n.needMore() == false {
			return nil
		}
		log.Debug("DialPeers", "peer", netAddr.String())
		peer, err := P2pComm.DialPeer(netAddr, &n.nodeInfo)
		if err != nil {
			log.Error("DialPeers", "Err", err.Error())
			continue
		}
		log.Debug("Addr perr", "peer", peer)
		n.AddPeer(peer)
		n.addrBook.AddAddress(netAddr)
		n.addrBook.Save()

	}
	return nil
}

func (n *Node) AddPeer(pr *peer) {
	n.omtx.Lock()
	defer n.omtx.Unlock()
	if pr.outbound == false {
		return
	}
	if peer, ok := n.outBound[pr.Addr()]; ok {
		log.Info("AddPeer", "delete peer", pr.Addr())
		n.addrBook.RemoveAddr(peer.Addr())
		n.addrBook.Save()
		delete(n.outBound, pr.Addr())
		peer.Close()
		peer = nil
	}
	log.Debug("AddPeer", "peer", pr.Addr())
	n.outBound[pr.Addr()] = pr
	pr.key = n.addrBook.key
	pr.Start()
	return

}

func (n *Node) Size() int {

	return n.nodeInfo.peerInfos.PeerSize()
}

func (n *Node) Has(paddr string) bool {
	n.omtx.Lock()
	defer n.omtx.Unlock()

	if _, ok := n.outBound[paddr]; ok {
		return true
	}
	return false
}

func (n *Node) GetRegisterPeer(paddr string) *peer {
	n.omtx.Lock()
	defer n.omtx.Unlock()
	if peer, ok := n.outBound[paddr]; ok {
		return peer
	}
	return nil
}

func (n *Node) GetRegisterPeers() []*peer {
	n.omtx.Lock()
	defer n.omtx.Unlock()
	var peers []*peer
	if len(n.outBound) == 0 {
		return peers
	}
	for _, peer := range n.outBound {
		peers = append(peers, peer)

	}
	return peers
}

func (n *Node) GetActivePeers() (map[string]*peer, map[string]*types.Peer) {
	regPeers := n.GetRegisterPeers()
	infos := n.nodeInfo.peerInfos.GetPeerInfos()

	var peers = make(map[string]*peer)
	for _, peer := range regPeers {
		if _, ok := infos[peer.Addr()]; ok {

			peers[peer.Addr()] = peer
		}
	}
	return peers, infos
}
func (n *Node) Remove(peerAddr string) {

	n.omtx.Lock()
	defer n.omtx.Unlock()
	peer, ok := n.outBound[peerAddr]
	if ok {
		delete(n.outBound, peerAddr)
		peer.Close()
		peer = nil
		return
	}

	return
}

func (n *Node) RemoveAll() {
	n.omtx.Lock()
	defer n.omtx.Unlock()
	for addr, peer := range n.outBound {
		delete(n.outBound, addr)
		peer.Close()
		peer = nil
	}
	return
}

func (n *Node) monitor() {
	go n.monitorErrPeer()
	go n.checkActivePeers()
	go n.getAddrFromOnline()
	go n.getAddrFromOffline()

}

func (n *Node) needMore() bool {
	outBoundNum := n.Size()
	if outBoundNum > MaxOutBoundNum || outBoundNum > StableBoundNum {
		return false
	}
	return true
}

func (n *Node) DetectNodeAddr() {

	var externalIp string
	for {
		cfg := n.nodeInfo.cfg
		LocalAddr = P2pComm.GetLocalAddr()
		log.Info("DetectNodeAddr", "addr:", LocalAddr)
		if len(LocalAddr) == 0 {
			log.Error("DetectNodeAddr", "NetWork Disable p2p Disable", "Retry until Network enable")
			time.Sleep(time.Second * 5)
			continue
		}
		if cfg.GetIsSeed() {
			log.Info("DetectNodeAddr", "ExIp", LocalAddr)
			externalIp = LocalAddr
			OutSide = true
			goto SET_ADDR
		}

		if len(cfg.Seeds) == 0 {
			goto SET_ADDR
		}
		for _, seed := range cfg.Seeds {

			pcli := NewP2pCli(nil)
			selfexaddrs, outside := pcli.GetExternIp(seed)
			if len(selfexaddrs) != 0 {
				OutSide = outside
				externalIp = selfexaddrs
				log.Info("DetectNodeAddr", " seed Exterip", externalIp)
				break
			}

		}
	SET_ADDR:
		//如果nat,getSelfExternalAddr 无法发现自己的外网地址，则把localaddr 赋值给外网地址
		if len(externalIp) == 0 {
			externalIp = LocalAddr
			log.Info("DetectNodeAddr", " SET_ADDR Exterip", externalIp)
		}
		rand.Seed(time.Now().Unix())
		exterPort := uint16(rand.Intn(64512) + 1023)
		addr := fmt.Sprintf("%v:%v", externalIp, exterPort)
		log.Debug("DetectionNodeAddr", "AddBlackList", addr)
		n.nodeInfo.blacklist.Add(addr) //把自己的外网地址加入到黑名单，以防连接self
		if exaddr, err := NewNetAddressString(addr); err == nil {
			n.nodeInfo.SetExternalAddr(exaddr)

		} else {
			log.Error("DetectionNodeAddr", "error", err.Error())
		}
		if listaddr, err := NewNetAddressString(fmt.Sprintf("%v:%v", LocalAddr, DefaultPort)); err == nil {
			n.nodeInfo.SetListenAddr(listaddr)
		}

		log.Info("DetectionNodeAddr", " Finish ExternalIp", externalIp, "LocalAddr", LocalAddr, "IsOutSide", OutSide)
		break
	}
}
