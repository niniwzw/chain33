package executor

import (
	"bytes"

	log "github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/common"
	pt "gitlab.33.cn/chain33/chain33/plugin/dapp/paracross/types"
	drivers "gitlab.33.cn/chain33/chain33/system/dapp"
	"gitlab.33.cn/chain33/chain33/types"
	"gitlab.33.cn/chain33/chain33/util"
)

var (
	clog                    = log.New("module", "execs.paracross")
	enableParacrossTransfer = true
)

type Paracross struct {
	drivers.DriverBase
}

func Init() {
	drivers.Register(GetName(), newParacross, 0)
	setPrefix()
	InitType()
}

func GetName() string {
	return newParacross().GetName()
}

func newParacross() drivers.Driver {
	c := &Paracross{}
	c.SetChild(c)
	return c
}

func (c *Paracross) GetName() string {
	return types.ParaX
}

func (c *Paracross) Exec(tx *types.Transaction, index int) (*types.Receipt, error) {
	_, err := c.DriverBase.Exec(tx, index)
	if err != nil {
		return nil, err
	}
	var payload pt.ParacrossAction
	err = types.Decode(tx.Payload, &payload)
	if err != nil {
		return nil, err
	}

	clog.Debug("Paracross.Exec", "payload.type", payload.Ty)
	if payload.Ty == pt.ParacrossActionCommit && payload.GetCommit() != nil {
		commit := payload.GetCommit()
		a := newAction(c, tx)
		return a.Commit(commit)
	} else if payload.Ty == pt.ParacrossActionTransfer && payload.GetAssetTransfer() != nil {
		clog.Debug("Paracross.Exec", "payload.type", payload.Ty, "transfer", "")
		_, err := c.checkTxGroup(tx, index)
		if err != nil {
			clog.Error("ParacrossActionTransfer", "get tx group failed", err, "hash", common.Bytes2Hex(tx.Hash()))
			return nil, err
		}
		a := newAction(c, tx)
		return a.AssetTransfer(payload.GetAssetTransfer())
	} else if payload.Ty == pt.ParacrossActionWithdraw && payload.GetAssetWithdraw() != nil {
		clog.Debug("Paracross.Exec", "payload.type", payload.Ty, "withdraw", "")
		_, err := c.checkTxGroup(tx, index)
		if err != nil {
			clog.Error("ParacrossActionWithdraw", "get tx group failed", err, "hash", common.Bytes2Hex(tx.Hash()))
			return nil, err
		}
		a := newAction(c, tx)
		return a.AssetWithdraw(payload.GetAssetWithdraw())
	} else if payload.Ty == pt.ParacrossActionMiner && payload.GetMiner() != nil {
		if index != 0 {
			return nil, types.ErrParaMinerBaseIndex
		}
		if !types.IsPara() {
			return nil, types.ErrNotSupport
		}
		a := newAction(c, tx)
		return a.Miner(payload.GetMiner())
	}

	return nil, types.ErrActionNotSupport
}

func (c *Paracross) checkTxGroup(tx *types.Transaction, index int) ([]*types.Transaction, error) {
	if tx.GroupCount >= 2 {
		txs, err := c.GetTxGroup(index)
		if err != nil {
			clog.Error("ParacrossActionTransfer", "get tx group failed", err, "hash", common.Bytes2Hex(tx.Hash()))
			return nil, err
		}
		return txs, nil
	}
	return nil, nil
}

//经para filter之后，交易组里面只会存在主链平行链跨链交易或全部平行链交易，全部平行链交易group里面有可能有资产转移交易
func crossTxGroupProc(txs []*types.Transaction, index int) ([]*types.Transaction, int32) {
	var headIdx, endIdx int32

	for i := index; i >= 0; i-- {
		if bytes.Equal(txs[index].Header, txs[i].Hash()) {
			headIdx = int32(i)
			break
		}
	}
	//cross mix tx, contain main and para tx, main prefix with types.ParaX
	endIdx = headIdx + txs[index].GroupCount
	for i := headIdx; i < endIdx; i++ {
		if bytes.HasPrefix(txs[i].Execer, []byte(types.ParaX)) {
			return txs[headIdx:endIdx], endIdx
		}
	}
	//cross asset transfer in tx group
	var transfers []*types.Transaction
	for i := headIdx; i < endIdx; i++ {
		if bytes.Contains(txs[i].Execer, []byte(types.ExecNamePrefix)) &&
			bytes.HasSuffix(txs[i].Execer, []byte(types.ParaX)) {
			transfers = append(transfers, txs[i])

		}
	}
	return transfers, endIdx

}

func (c *Paracross) ExecLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := c.DriverBase.ExecLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	//执行成功
	var payload pt.ParacrossAction
	err = types.Decode(tx.Payload, &payload)
	if err != nil {
		return nil, err
	}

	for _, log := range receipt.Logs {
		if log.Ty == pt.TyLogParacrossCommit {
			var g pt.ReceiptParacrossCommit
			types.Decode(log.Log, &g)

			var r pt.ParacrossTx
			r.TxHash = string(tx.Hash())
			set.KV = append(set.KV, &types.KeyValue{calcLocalTxKey(g.Status.Title, g.Status.Height, tx.From()), types.Encode(&r)})
		} else if log.Ty == pt.TyLogParacrossCommitDone {
			var g pt.ReceiptParacrossDone
			types.Decode(log.Log, &g)

			key := calcLocalTitleKey(g.Title)
			set.KV = append(set.KV, &types.KeyValue{key, types.Encode(&g)})

			key = calcTitleHeightKey(g.Title, g.Height)
			set.KV = append(set.KV, &types.KeyValue{key, types.Encode(&g)})

			r, err := c.saveLocalParaTxs(tx, false)
			if err != nil {
				return nil, err
			}
			set.KV = append(set.KV, r.KV...)
		} else if log.Ty == pt.TyLogParacrossCommitRecord {
			var g pt.ReceiptParacrossRecord
			types.Decode(log.Log, &g)

			var r pt.ParacrossTx
			r.TxHash = string(tx.Hash())
			set.KV = append(set.KV, &types.KeyValue{calcLocalTxKey(g.Status.Title, g.Status.Height, tx.From()), types.Encode(&r)})
		} else if log.Ty == pt.TyLogParacrossMiner {
			if index != 0 {
				return nil, types.ErrParaMinerBaseIndex
			}
			var g pt.ReceiptParacrossMiner
			types.Decode(log.Log, &g)

			var mixTxHashs, paraTxHashs, crossTxHashs [][]byte
			txs := c.GetTxs()
			//remove the 0 vote tx
			for i := 1; i < len(txs); i++ {
				tx := txs[i]
				hash := tx.Hash()
				mixTxHashs = append(mixTxHashs, hash)
				//跨链交易包含了主链交易，需要过滤出来
				if bytes.Contains(tx.Execer, []byte(types.ExecNamePrefix)) {
					paraTxHashs = append(paraTxHashs, hash)
				}
			}
			for i := 1; i < len(txs); i++ {
				tx := txs[i]
				if tx.GroupCount >= 2 {
					crossTxs, end := crossTxGroupProc(txs, i)
					for _, crossTx := range crossTxs {
						crossTxHashs = append(crossTxHashs, crossTx.Hash())
					}
					i = int(end) - 1
					continue
				}
				if bytes.Contains(tx.Execer, []byte(types.ExecNamePrefix)) &&
					bytes.HasSuffix(tx.Execer, []byte(types.ParaX)) {
					crossTxHashs = append(crossTxHashs, tx.Hash())
				}
			}

			g.Status.TxHashs = paraTxHashs
			g.Status.TxResult = util.CalcSubBitMap(mixTxHashs, paraTxHashs, c.GetReceipt()[1:])
			g.Status.CrossTxHashs = crossTxHashs
			g.Status.CrossTxResult = util.CalcSubBitMap(mixTxHashs, crossTxHashs, c.GetReceipt()[1:])

			set.KV = append(set.KV, &types.KeyValue{pt.CalcMinerHeightKey(g.Status.Title, g.Status.Height), types.Encode(g.Status)})

		} else if log.Ty == pt.TyLogParaAssetWithdraw {
			// TODO
		} else if log.Ty == pt.TyLogParaAssetTransfer {
			// TODO
		} else if log.Ty == types.TyLogExecTransfer {
			//  主链转出记录，
			//  转入在 commit done 时几率， 因为没有日志里没有当时tx信息
			var g pt.ParacrossAction
			err = types.Decode(tx.Payload, &g)
			if g.Ty == pt.ParacrossActionTransfer {
				r, err := c.initLocalAssetTransfer(tx, true, false)
				if err != nil {
					return nil, err
				}
				set.KV = append(set.KV, r)

			} // else if tx.ActionName() == pt.ParacrossActionCommitStr {
			//
			//}

		} /* Token 暂时不支持 */
	}
	return set, nil
}

func (c *Paracross) saveLocalParaTxs(tx *types.Transaction, isDel bool) (*types.LocalDBSet, error) {
	var set types.LocalDBSet

	var payload pt.ParacrossAction
	err := types.Decode(tx.Payload, &payload)
	if err != nil {
		return nil, err
	}
	if payload.Ty != pt.ParacrossActionCommit || payload.GetCommit() == nil {
		return nil, nil
	}

	commit := payload.GetCommit()
	for i := 0; i < len(commit.Status.CrossTxHashs); i++ {
		success := util.BitMapBit(commit.Status.CrossTxResult, uint32(i))

		paraTx, err := GetTx(c.GetApi(), commit.Status.CrossTxHashs[i])
		if err != nil {
			clog.Crit("paracross.Commit Load Tx failed", "para title", commit.Status.Title,
				"para height", commit.Status.Height, "para tx index", i, "error", err, "txHash",
				common.Bytes2Hex(commit.Status.CrossTxHashs[i]))
			return nil, err
		}

		var payload pt.ParacrossAction
		err = types.Decode(paraTx.Tx.Payload, &payload)
		if err != nil {
			clog.Crit("paracross.Commit Decode Tx failed", "para title", commit.Status.Title,
				"para height", commit.Status.Height, "para tx index", i, "error", err, "txHash",
				common.Bytes2Hex(commit.Status.CrossTxHashs[i]))
			return nil, err
		}
		if payload.Ty == pt.ParacrossActionTransfer {
			kv, err := c.updateLocalAssetTransfer(tx, paraTx.Tx, success, isDel)
			if err != nil {
				return nil, err
			}
			set.KV = append(set.KV, kv)
		} else if payload.Ty == pt.ParacrossActionWithdraw {
			kv, err := c.initLocalAssetWithdraw(tx, paraTx.Tx, true, success, isDel)
			if err != nil {
				return nil, err
			}
			set.KV = append(set.KV, kv)
		}
	}

	return &set, nil
}

func getCommitHeight(payload []byte) (int64, error) {
	var a pt.ParacrossAction
	err := types.Decode(payload, &a)
	if err != nil {
		return 0, err
	}
	if a.GetCommit() == nil {
		return 0, types.ErrInputPara
	}

	return a.GetCommit().Status.Height, nil
}

func (c *Paracross) initLocalAssetTransfer(tx *types.Transaction, success, isDel bool) (*types.KeyValue, error) {
	clog.Info("para execLocal", "tx hash", common.Bytes2Hex(tx.Hash()))
	clog.Info("para execLocal 1", "action name", tx.ActionName())
	key := calcLocalAssetKey(tx.Hash())
	if isDel {
		c.GetLocalDB().Set(key, nil)
		return &types.KeyValue{key, nil}, nil
	}

	var asset pt.ParacrossAsset
	clog.Info("para execLocal 2", "action name", tx.ActionName())
	amount, err := tx.Amount()
	if err != nil {
		return nil, err
	}

	asset = pt.ParacrossAsset{
		From:       tx.From(),
		To:         tx.To,
		Amount:     amount,
		IsWithdraw: false,
		TxHash:     tx.Hash(),
		Height:     c.GetHeight(),
	}
	clog.Info("para execLocal 3", "action name", tx.ActionName())

	err = c.GetLocalDB().Set(key, types.Encode(&asset))
	if err != nil {
		clog.Error("para execLocal", "set", common.Bytes2Hex(tx.Hash()), "failed", err)
	}
	return &types.KeyValue{key, types.Encode(&asset)}, nil
}

func (c *Paracross) initLocalAssetWithdraw(txCommit, tx *types.Transaction, isWithdraw, success, isDel bool) (*types.KeyValue, error) {
	key := calcLocalAssetKey(tx.Hash())
	if isDel {
		c.GetLocalDB().Set(key, nil)
		return &types.KeyValue{key, nil}, nil
	}

	var asset pt.ParacrossAsset

	amount, err := tx.Amount()
	if err != nil {
		return nil, err
	}
	asset.ParaHeight, err = getCommitHeight(txCommit.Payload)
	if err != nil {
		return nil, err
	}
	asset = pt.ParacrossAsset{
		From:       tx.From(),
		To:         tx.To,
		Amount:     amount,
		IsWithdraw: isWithdraw,
		TxHash:     tx.Hash(),
		Height:     c.GetHeight(),
	}

	asset.CommitDoneHeight = c.GetHeight()
	asset.Success = success

	err = c.GetLocalDB().Set(key, types.Encode(&asset))
	if err != nil {
		clog.Error("para execLocal", "set", "", "failed", err)
	}
	return &types.KeyValue{key, types.Encode(&asset)}, nil
}

func (c *Paracross) updateLocalAssetTransfer(txCommit, tx *types.Transaction, success, isDel bool) (*types.KeyValue, error) {
	clog.Info("para execLocal", "tx hash", common.Bytes2Hex(tx.Hash()))
	key := calcLocalAssetKey(tx.Hash())

	var asset pt.ParacrossAsset
	v, err := c.GetLocalDB().Get(key)
	if err != nil {
		return nil, err
	}
	err = types.Decode(v, &asset)
	if err != nil {
		panic(err)
	}
	if !isDel {
		asset.ParaHeight, err = getCommitHeight(txCommit.Payload)
		if err != nil {
			return nil, err
		}
		asset.CommitDoneHeight = c.GetHeight()
		asset.Success = success
	} else {
		asset.ParaHeight = 0
		asset.CommitDoneHeight = 0
		asset.Success = false
	}
	c.GetLocalDB().Set(key, types.Encode(&asset))
	return &types.KeyValue{key, types.Encode(&asset)}, nil
}

func (c *Paracross) ExecDelLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := c.DriverBase.ExecDelLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	//执行成功
	//执行成功
	var payload pt.ParacrossAction
	err = types.Decode(tx.Payload, &payload)
	if err != nil {
		return nil, err
	}

	for _, log := range receipt.Logs {
		if log.Ty == pt.TyLogParacrossCommit { //} || log.Ty == types.TyLogParacrossCommitRecord {
			var g pt.ReceiptParacrossCommit
			types.Decode(log.Log, &g)

			var r pt.ParacrossTx
			r.TxHash = string(tx.Hash())
			set.KV = append(set.KV, &types.KeyValue{calcLocalTxKey(g.Status.Title, g.Status.Height, tx.From()), nil})
		} else if log.Ty == pt.TyLogParacrossCommitDone {
			var g pt.ReceiptParacrossDone
			types.Decode(log.Log, &g)
			g.Height = g.Height - 1

			key := calcLocalTitleKey(g.Title)
			set.KV = append(set.KV, &types.KeyValue{key, types.Encode(&g)})

			key = calcTitleHeightKey(g.Title, g.Height)
			set.KV = append(set.KV, &types.KeyValue{key, nil})

			r, err := c.saveLocalParaTxs(tx, true)
			if err != nil {
				return nil, err
			}
			set.KV = append(set.KV, r.KV...)
		} else if log.Ty == pt.TyLogParacrossCommitRecord {
			var g pt.ReceiptParacrossRecord
			types.Decode(log.Log, &g)

			var r pt.ParacrossTx
			r.TxHash = string(tx.Hash())
			set.KV = append(set.KV, &types.KeyValue{calcLocalTxKey(g.Status.Title, g.Status.Height, tx.From()), nil})
		} else if log.Ty == pt.TyLogParacrossMiner {
			var g pt.ReceiptParacrossMiner
			types.Decode(log.Log, &g)

			set.KV = append(set.KV, &types.KeyValue{pt.CalcMinerHeightKey(g.Status.Title, g.Status.Height), nil})

		} else if log.Ty == pt.TyLogParaAssetWithdraw {
			// TODO
		} else if log.Ty == pt.TyLogParaAssetTransfer {
			// TODO
		} else if log.Ty == types.TyLogExecTransfer {
			//  主链转出记录，
			//  转入在 commit done 时几率， 因为没有日志里没有当时tx信息
			var g pt.ParacrossAction
			err = types.Decode(tx.Payload, &g)
			if g.Ty == pt.ParacrossActionTransfer {
				r, err := c.initLocalAssetTransfer(tx, true, true)
				if err != nil {
					return nil, err
				}
				set.KV = append(set.KV, r)

			} // else if tx.ActionName() == pt.ParacrossActionCommitStr {
			//
			//}

		} /* Token 暂时不支持 */
	}
	return set, nil
}

func (c *Paracross) Query(funcName string, params []byte) (types.Message, error) {
	if funcName == "ParacrossGetTitle" {
		var in types.ReqStr
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.ParacrossGetHeight(in.ReqStr)
	} else if funcName == "ParacrossListTitles" {
		return c.ParacrossListTitles()
	} else if funcName == "ParacrossGetTitleHeight" {
		var in pt.ReqParacrossTitleHeight
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.ParacrossGetTitleHeight(in.Title, in.Height)
	} else if funcName == "ParacrossGetAssetTxResult" {
		var in types.ReqHash
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.ParacrossGetAssetTxResult(in.Hash)
	}

	return nil, types.ErrActionNotSupport
}

func (c *Paracross) IsFriend(myexec, writekey []byte, tx *types.Transaction) bool {
	//不允许平行链
	if types.IsPara() {
		return false
	}
	//只允许写名字为主链名称的
	if string(myexec) != c.GetName() {
		return false
	}
	//只允许跨链交易
	return c.Allow(tx, 0) == nil
}

func (c *Paracross) Allow(tx *types.Transaction, index int) error {
	err := c.DriverBase.Allow(tx, index)
	if err == nil {
		return nil
	}
	// 增加新的规则: 在主链执行器带着title的 asset-transfer/asset-withdraw 交易允许执行
	// 1. user.p.${tilte}.${paraX}
	// 1. payload 的 actionType = t/w
	if !types.IsPara() && c.allowIsParaAssetTx(tx.Execer) {
		var payload pt.ParacrossAction
		err = types.Decode(tx.Payload, &payload)
		if err != nil {
			return err
		}
		if payload.Ty == pt.ParacrossActionTransfer || payload.Ty == pt.ParacrossActionWithdraw {
			return nil
		}
	}
	return types.ErrNotAllow
}

func (c *Paracross) allowIsParaAssetTx(execer []byte) bool {
	if !bytes.HasPrefix(execer, types.ParaKey) {
		return false
	}
	count := 0
	index := 0
	s := len(types.ParaKey)
	for i := s; i < len(execer); i++ {
		if execer[i] == '.' {
			count++
			index = i
		}
	}
	if count == 1 && c.AllowIsSame(execer[index+1:]) {
		return true
	}
	return false
}
