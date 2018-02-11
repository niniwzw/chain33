package p2p

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"code.aliyun.com/chain33/chain33/p2p/nat"
	pb "code.aliyun.com/chain33/chain33/types"

	"google.golang.org/grpc"
)

func (l *DefaultListener) Close() bool {

	l.listener.Close()
	l.server.Stop()
	log.Info("stop", "DefaultListener", "close")
	close(l.natClose)
	log.Info("stop", "NatMapPort", "close")
	close(l.nodeInfo.p2pBroadcastChan) //close p2pserver manageStream
	close(l.p2pserver.loopdone)
	return true
}

type Listener interface {
	Close() bool
}

// Implements Listener
type DefaultListener struct {
	listener  net.Listener
	server    *grpc.Server
	nodeInfo  *NodeInfo
	p2pserver *p2pServer
	natClose  chan struct{}
	n         *Node
}

func NewDefaultListener(protocol string, node *Node) Listener {

	log.Debug("NewDefaultListener", "localPort", DefaultPort)
	listener, err := net.Listen(protocol, fmt.Sprintf(":%v", DefaultPort))
	if err != nil {
		log.Crit("Failed to listen", "Error", err.Error())
		return nil
	}

	dl := &DefaultListener{
		listener: listener,
		nodeInfo: node.nodeInfo,
		natClose: make(chan struct{}, 1),
		n:        node,
	}
	pServer := NewP2pServer()
	pServer.node = dl.n
	pServer.ManageStream()
	dl.server = grpc.NewServer()
	dl.p2pserver = pServer
	pb.RegisterP2PgserviceServer(dl.server, pServer)
	go dl.NatMapPort()
	go dl.listenRoutine()
	return dl
}

func (l *DefaultListener) WaitForNat() {
	<-l.nodeInfo.natNoticeChain
	return
}

func (l *DefaultListener) NatMapPort() {
	if OutSide == true { //在外网的节点不需要映射端口
		return
	}
	l.WaitForNat()
	var err error
	for i := 0; i < TryMapPortTimes; i++ {

		err = nat.Any().AddMapping("TCP", int(l.n.nodeInfo.GetExternalAddr().Port), int(DefaultPort), "chain33-p2p", time.Minute*20)
		if err != nil {
			log.Error("NatMapPort", "err", err.Error())
			l.n.FlushNodePort(DefaultPort, uint16(rand.Intn(64512)+1023))
			log.Info("NatMapPort", "New External Port", l.n.nodeInfo.GetExternalAddr())
			continue
		}

		break
	}
	if err != nil {
		//映射失败
		log.Warn("NatMapPort", "Nat Faild", "Sevice=6")
		l.nodeInfo.natResultChain <- false
		return
	}

	l.nodeInfo.natResultChain <- true
	refresh := time.NewTimer(mapUpdateInterval)
	for {
		select {
		case <-refresh.C:
			nat.Any().AddMapping("TCP", int(l.n.nodeInfo.GetExternalAddr().Port), int(DefaultPort), "chain33-p2p", time.Minute*20)
			refresh.Reset(mapUpdateInterval)

		case <-l.natClose:
			nat.Any().DeleteMapping("TCP", int(l.n.nodeInfo.GetExternalAddr().Port), int(DefaultPort))
			return
		}
	}

}

func (l *DefaultListener) listenRoutine() {
	log.Debug("LISTENING", "Start Listening+++++++++++++++++Port", l.nodeInfo.listenAddr.Port)
	l.server.Serve(l.listener)
}
