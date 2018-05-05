package p2p

import (
	"fmt"
	"net"
	"time"

	pb "gitlab.33.cn/chain33/chain33/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Listener interface {
	Close()
	Start()
}

func (l *listener) Start() {
	l.p2pserver.Start()
	go l.server.Serve(l.netlistener)

}
func (l *listener) Close() {
	l.netlistener.Close()
	go l.server.Stop()
	l.p2pserver.Close()
	log.Info("stop", "listener", "close")

}

type listener struct {
	server      *grpc.Server
	nodeInfo    *NodeInfo
	p2pserver   *P2pServer
	node        *Node
	netlistener net.Listener
}

func NewListener(protocol string, node *Node) Listener {
	log.Debug("NewListener", "localPort", defaultPort)
	l, err := net.Listen(protocol, fmt.Sprintf(":%v", defaultPort))
	if err != nil {
		log.Crit("Failed to listen", "Error", err.Error())
		return nil
	}

	dl := &listener{
		nodeInfo:    node.nodeInfo,
		node:        node,
		netlistener: l,
	}

	pServer := NewP2pServer()
	pServer.node = dl.node

	//区块最多10M
	msgRecvOp := grpc.MaxMsgSize(11 * 1024 * 1024)     //设置最大接收数据大小位11M
	msgSendOp := grpc.MaxSendMsgSize(11 * 1024 * 1024) //设置最大发送数据大小为11M

	var keepparm keepalive.ServerParameters
	keepparm.Time = 3 * time.Minute
	keepparm.Timeout = 30 * time.Second
	keepparm.MaxConnectionIdle = 5 * time.Second
	maxStreams := grpc.MaxConcurrentStreams(1000)
	keepOp := grpc.KeepaliveParams(keepparm)
	connTimeout := grpc.ConnectionTimeout(0)
	dl.server = grpc.NewServer(msgRecvOp, msgSendOp, keepOp, maxStreams, connTimeout)
	dl.p2pserver = pServer
	pb.RegisterP2PgserviceServer(dl.server, pServer)
	return dl
}
