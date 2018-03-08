package none

//package none execer for unknow execer
//all none transaction exec ok, execept nofee
//nofee transaction will not pack into block

import (
	"code.aliyun.com/chain33/chain33/executor/drivers"
)

func init() {
	n := newNone()
	drivers.Register(n.GetName(), n)
}

type None struct {
	drivers.DriverBase
}

func newNone() *None {
	n := &None{}
	n.SetChild(n)
	return n
}

func (n *None) GetName() string {
	return "none"
}
