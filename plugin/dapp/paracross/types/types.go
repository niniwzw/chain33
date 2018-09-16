package types

import (
	fmt "fmt"
	"math/rand"
	"time"

	"github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/common/address"
	"gitlab.33.cn/chain33/chain33/types"
)

const ParaCrossX = "paracross"

var tlog = log15.New("module", ParaCrossX)

const (
	// paracross 执行器的日志类型
	TyLogParacrossCommit     = 650
	TyLogParacrossCommitDone = 651
	// record 和 commit 不一样， 对应高度完成共识后收到commit 交易
	// 这个交易就不参与共识, 只做记录
	TyLogParacrossCommitRecord = 652
	TyLogParaAssetTransfer     = 653
	TyLogParaAssetWithdraw     = 654
	//在平行链上保存节点参与共识的数据
	TyLogParacrossMiner = 655
)

type ParacrossCommitTx struct {
	Fee    int64               `json:"fee"`
	Status ParacrossNodeStatus `json:"status"`
}

// action type
const (
	ParacrossActionCommit = iota
	ParacrossActionMiner
)
const (
	ParacrossActionTransfer = iota + types.ParaCrossTransferActionTypeStart
	ParacrossActionWithdraw
)

// status
const (
	ParacrossStatusCommiting = iota
	ParacrossStatusCommitDone
)

var (
	ParacrossActionCommitStr   = string("Commit")
	ParacrossTransferPerfix    = "crossPara."
	ParacrossActionTransferStr = ParacrossTransferPerfix + string("Transfer")
	ParacrossActionWithdrawStr = ParacrossTransferPerfix + string("Withdraw")
)

func CalcMinerHeightKey(title string, height int64) []byte {
	paraVoteHeightKey := types.ExecName("paracross") + "-titleVoteHeight-"
	return []byte(fmt.Sprintf(paraVoteHeightKey+"%s-%012d", title, height))
}

func CreateRawCommitTx4MainChain(status *ParacrossNodeStatus, name string, fee int64) (*types.Transaction, error) {
	return createRawCommitTx(status, name, fee)
}

func CreateRawParacrossCommitTx(parm *ParacrossCommitTx) (*types.Transaction, error) {
	if parm == nil {
		tlog.Error("CreateRawParacrossCommitTx", "parm", parm)
		return nil, types.ErrInvalidParam
	}
	return createRawCommitTx(&parm.Status, types.ExecName(ParaCrossX), parm.Fee)
}

func createRawCommitTx(status *ParacrossNodeStatus, name string, fee int64) (*types.Transaction, error) {
	v := &ParacrossCommitAction{
		Status: status,
	}
	action := &ParacrossAction{
		Ty:    ParacrossActionCommit,
		Value: &ParacrossAction_Commit{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(name),
		Payload: types.Encode(action),
		Fee:     fee,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(name),
	}

	err := tx.SetRealFee(types.MinFee)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CreateRawTransferTx(param *types.CreateTx) (*types.Transaction, error) {
	// 跨链交易需要在主链和平行链上执行， 所以应该可以在主链和平行链上构建
	if !types.IsParaExecName(param.GetExecName()) {
		tlog.Error("CreateRawTransferTx", "exec", param.GetExecName())
		return nil, types.ErrInputPara
	}

	transfer := &ParacrossAction{}
	if !param.IsWithdraw {
		v := &ParacrossAction_AssetTransfer{AssetTransfer: &types.AssetsTransfer{
			Amount: param.Amount, Note: param.GetNote(), To: param.GetTo()}}
		transfer.Value = v
		transfer.Ty = ParacrossActionTransfer
	} else {
		v := &ParacrossAction_AssetWithdraw{AssetWithdraw: &types.AssetsWithdraw{
			Amount: param.Amount, Note: param.GetNote(), To: param.GetTo()}}
		transfer.Value = v
		transfer.Ty = ParacrossActionWithdraw
	}
	tx := &types.Transaction{
		Execer:  []byte(param.GetExecName()),
		Payload: types.Encode(transfer),
		To:      address.ExecAddress(param.GetExecName()),
		Fee:     param.Fee,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
	}

	if err := tx.SetRealFee(types.MinFee); err != nil {
		return nil, err
	}

	return tx, nil
}

func CreateRawMinerTx(status *ParacrossNodeStatus) (*types.Transaction, error) {
	v := &ParacrossMinerAction{
		Status: status,
	}
	action := &ParacrossAction{
		Ty:    ParacrossActionMiner,
		Value: &ParacrossAction_Miner{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(types.ExecName(ParaCrossX)),
		Payload: types.Encode(action),
		Nonce:   0, //for consensus purpose, block hash need same, different auth node need keep totally same vote tx
		To:      address.ExecAddress(types.ExecName(ParaCrossX)),
	}

	err := tx.SetRealFee(types.MinFee)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
