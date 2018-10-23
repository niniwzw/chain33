package commands

import (
	//	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.33.cn/chain33/chain33/common"
	"gitlab.33.cn/chain33/chain33/common/address"
	"gitlab.33.cn/chain33/chain33/rpc/jsonclient"
	rpctypes "gitlab.33.cn/chain33/chain33/rpc/types"
	cty "gitlab.33.cn/chain33/chain33/system/dapp/coins/types"
	"gitlab.33.cn/chain33/chain33/types"

	// TODO: 暂时将插件中的类型引用起来，后续需要修改

	"encoding/hex"
	"math"
	"math/rand"
	"time"
)

func decodeTransaction(tx *rpctypes.Transaction) *TxResult {
	feeResult := strconv.FormatFloat(float64(tx.Fee)/float64(types.Coin), 'f', 4, 64)
	result := &TxResult{
		Execer:     tx.Execer,
		Payload:    tx.Payload,
		RawPayload: tx.RawPayload,
		Signature:  tx.Signature,
		Fee:        feeResult,
		Expire:     tx.Expire,
		Nonce:      tx.Nonce,
		To:         tx.To,
		From:       tx.From,
		GroupCount: tx.GroupCount,
		Header:     tx.Header,
		Next:       tx.Next,
	}

	if tx.Amount != 0 {
		result.Amount = strconv.FormatFloat(float64(tx.Amount)/float64(types.Coin), 'f', 4, 64)
	}
	if tx.From != "" {
		result.From = tx.From
	}
	payload, err := common.FromHex(tx.RawPayload)
	if err != nil {
		return result
	}
	exec := types.LoadExecutorType(tx.Execer)
	if exec != nil {
		tx.Payload, _ = exec.DecodePayload(&types.Transaction{Payload: payload})
	}
	return result
}

func decodeAccount(acc *types.Account, precision int64) *AccountResult {
	balanceResult := strconv.FormatFloat(float64(acc.GetBalance())/float64(precision), 'f', 4, 64)
	frozenResult := strconv.FormatFloat(float64(acc.GetFrozen())/float64(precision), 'f', 4, 64)
	accResult := &AccountResult{
		Addr:     acc.GetAddr(),
		Currency: acc.GetCurrency(),
		Balance:  balanceResult,
		Frozen:   frozenResult,
	}
	return accResult
}

func decodeLog(execer []byte, rlog rpctypes.ReceiptDataResult) *ReceiptData {
	rd := &ReceiptData{Ty: rlog.Ty, TyName: rlog.TyName}
	for _, l := range rlog.Logs {
		rl := &ReceiptLog{Ty: l.Ty, TyName: l.TyName, RawLog: l.RawLog}
		logty := types.LoadLog(execer, int64(l.Ty))
		if logty == nil {
			rl.Log = nil
		}
		data, err := common.FromHex(l.RawLog)
		if err != nil {
			rl.Log = nil
		}
		rl.Log, err = logty.Json(data)
		if err != nil {
			rl.Log = nil
		}
		rd.Logs = append(rd.Logs, rl)
	}
	return rd
}

func SendToAddress(rpcAddr string, from string, to string, amount int64, note string, isToken bool, tokenSymbol string, isWithdraw bool) {
	amt := amount
	if isWithdraw {
		amt = -amount
	}
	params := types.ReqWalletSendToAddress{From: from, To: to, Amount: amt, Note: note}
	if !isToken {
		params.IsToken = false
	} else {
		params.IsToken = true
		params.TokenSymbol = tokenSymbol
	}

	var res rpctypes.ReplyHash
	ctx := jsonclient.NewRpcCtx(rpcAddr, "Chain33.SendToAddress", params, &res)
	ctx.Run()
}

func CreateRawTx(cmd *cobra.Command, to string, amount float64, note string, isWithdraw bool, tokenSymbol, execName string) (string, error) {
	if amount < 0 {
		return "", types.ErrAmount
	}
	paraName, _ := cmd.Flags().GetString("paraName")
	amountInt64 := int64(math.Trunc((amount+0.0000001)*1e4)) * 1e4
	initExecName := execName
	execName = getRealExecName(paraName, execName)
	if execName != "" && !types.IsAllowExecName([]byte(execName), []byte(execName)) {
		return "", types.ErrExecNameNotMatch
	}
	var tx *types.Transaction
	transfer := &cty.CoinsAction{}
	if !isWithdraw {
		if initExecName != "" {
			v := &cty.CoinsAction_TransferToExec{TransferToExec: &types.AssetsTransferToExec{Amount: amountInt64, Note: note, ExecName: execName}}
			transfer.Value = v
			transfer.Ty = cty.CoinsActionTransferToExec
		} else {
			v := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: amountInt64, Note: note, To: to}}
			transfer.Value = v
			transfer.Ty = cty.CoinsActionTransfer
		}
	} else {
		v := &cty.CoinsAction_Withdraw{Withdraw: &types.AssetsWithdraw{Amount: amountInt64, Note: note, ExecName: execName}}
		transfer.Value = v
		transfer.Ty = cty.CoinsActionWithdraw
	}
	execer := []byte(getRealExecName(paraName, "coins"))
	if paraName == "" {
		tx = &types.Transaction{Execer: execer, Payload: types.Encode(transfer), To: to}
	} else {
		tx = &types.Transaction{Execer: execer, Payload: types.Encode(transfer), To: address.ExecAddress(string(execer))}
	}

	var err error
	tx.Fee, err = tx.GetRealFee(types.MinFee)
	if err != nil {
		return "", err
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	txHex := types.Encode(tx)
	return hex.EncodeToString(txHex), nil
}

func GetExecAddr(exec string) (string, error) {
	if ok := types.IsAllowExecName([]byte(exec), []byte(exec)); !ok {
		return "", types.ErrExecNameNotAllow
	}

	addrResult := address.ExecAddress(exec)
	result := addrResult
	return result, nil
}

func getExecuterNameString() string {
	str := "executer name (only "
	allowExeName := types.AllowUserExec
	nameLen := len(allowExeName)
	for i := 0; i < nameLen; i++ {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("\"%s\"", string(allowExeName[i]))
	}
	str += " and user-defined type supported)"
	return str
}

// FormatAmountValue2Display 将传输、计算的amount值格式化成显示值
func FormatAmountValue2Display(amount int64) string {
	return strconv.FormatFloat(float64(amount)/float64(types.Coin), 'f', 4, 64)
}

// FormatAmountDisplay2Value 将显示、输入的amount值格式话成传输、计算值
func FormatAmountDisplay2Value(amount float64) int64 {
	return int64(amount*types.InputPrecision) * types.Multiple1E4
}

// GetAmountValue 将命令行中的amount值转换成int64
func GetAmountValue(cmd *cobra.Command, field string) int64 {
	amount, _ := cmd.Flags().GetFloat64(field)
	return FormatAmountDisplay2Value(amount)
}

func getRealExecName(paraName string, name string) string {
	if strings.HasPrefix(name, "user.p.") {
		return name
	}
	return paraName + name
}
