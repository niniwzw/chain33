package rpc

import (
	"golang.org/x/net/context"

	pb "code.aliyun.com/chain33/chain33/types"
)

type Grpc struct {
	gserver *grpcServer
}

func (req *Grpc) SendTransaction(ctx context.Context, in *pb.Transaction) (*pb.Reply, error) {

	cli := NewClient("channel", "")
	cli.SetQueue(req.gserver.q)

	reply := cli.SendTx(in)
	return &pb.Reply{IsOk: true, Msg: reply.GetData().(*pb.Reply).Msg}, reply.Err()

}

func (req *Grpc) QueryTransaction(ctx context.Context, in *pb.ReqHash) (*pb.Reply, error) {
	cli := NewClient("channel", "")
	cli.SetQueue(req.gserver.q)

	reply, err := cli.QueryTx(in.Hash)
	if err != nil {
		return nil, err
	}

	return &pb.Reply{IsOk: true, Msg: (interface{}(reply)).(*pb.Reply).Msg}, nil

}

func (req *Grpc) GetBlocks(ctx context.Context, in *pb.ReqBlocks) (*pb.Reply, error) {

	cli := NewClient("channel", "")
	cli.SetQueue(req.gserver.q)

	reply, err := cli.GetBlocks(in.Start, in.End, in.Isdetail)
	if err != nil {
		return nil, err
	}

	return &pb.Reply{IsOk: true, Msg: (interface{}(reply)).(*pb.Reply).Msg}, nil

}
