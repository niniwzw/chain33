package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	log "github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/common"
	"gitlab.33.cn/chain33/chain33/common/address"
	"gitlab.33.cn/chain33/chain33/types/jsonpb"

	_ "gitlab.33.cn/chain33/chain33/common/crypto/ecdsa"
	_ "gitlab.33.cn/chain33/chain33/common/crypto/ed25519"
	_ "gitlab.33.cn/chain33/chain33/common/crypto/secp256k1"
	_ "gitlab.33.cn/chain33/chain33/common/crypto/sm2"
)

var tlog = log.New("module", "types")

const Size_1K_shiftlen uint = 10

type Message proto.Message

type Query4Cli struct {
	Execer   string      `json:"execer"`
	FuncName string      `json:"funcName"`
	Payload  interface{} `json:"payload"`
}

//交易组的接口，Transactions 和 Transaction 都符合这个接口
type TxGroup interface {
	Tx() *Transaction
	GetTxGroup() (*Transactions, error)
	CheckSign() bool
}

func ExecName(name string) string {
	if IsParaExecName(name) {
		return name
	}
	return ExecNamePrefix + name
}

func IsAllowExecName(name []byte, execer []byte) bool {
	// name长度不能超过系统限制
	if len(name) > address.MaxExecNameLength {
		return false
	}
	// name中不允许有 "-"
	if bytes.Contains(name, slash) {
		return false
	}
	if !bytes.Equal(name, execer) && !bytes.Equal(name, GetRealExecName(execer)) {
		return false
	}
	if bytes.HasPrefix(name, UserKey) {
		return true
	}
	for i := range AllowUserExec {
		if bytes.Equal(AllowUserExec[i], name) {
			return true
		}
	}
	return false
}

var bytesExec = []byte("exec-")
var commonPrefix = []byte("mavl-")

func GetExecKey(key []byte) (string, bool) {
	n := 0
	start := 0
	end := 0
	for i := len(commonPrefix); i < len(key); i++ {
		if key[i] == '-' {
			n = n + 1
			if n == 2 {
				start = i + 1
			}
			if n == 3 {
				end = i
				break
			}
		}
	}
	if start > 0 && end > 0 {
		if bytes.Equal(key[start:end+1], bytesExec) {
			//find addr
			start = end + 1
			for k := end; k < len(key); k++ {
				if key[k] == ':' { //end+1
					end = k
					return string(key[start:end]), true
				}
			}
		}
	}
	return "", false
}

func FindExecer(key []byte) (execer []byte, err error) {
	if !bytes.HasPrefix(key, commonPrefix) {
		return nil, ErrMavlKeyNotStartWithMavl
	}
	for i := len(commonPrefix); i < len(key); i++ {
		if key[i] == '-' {
			return key[len(commonPrefix):i], nil
		}
	}
	return nil, ErrNoExecerInMavlKey
}

func GetRealExecName(execer []byte) []byte {
	//平行链执行器，获取真实执行器的规则
	if bytes.HasPrefix(execer, ParaKey) {
		count := 0
		for i := 0; i < len(execer); i++ {
			if execer[i] == '.' {
				count++
			}
			if count == 3 && i < (len(execer)-1) {
				newexec := execer[i+1:]
				if bytes.HasPrefix(newexec, UserKey) && !bytes.HasPrefix(newexec, ParaKey) {
					return GetRealExecName(newexec)
				}
				return newexec
			}
		}
	} else if bytes.HasPrefix(execer, UserKey) {
		//不是user.p. 的情况, 而是user. 的情况
		count := 0
		index := 0
		for i := 0; i < len(execer); i++ {
			if execer[i] == '.' {
				count++
			}
			index = i
			if count == 2 {
				index--
				break
			}
		}
		e := execer[5 : index+1]
		if len(e) > 0 {
			return e
		}
	}
	return execer
}

func Encode(data proto.Message) []byte {
	b, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}
	return b
}

func Size(data proto.Message) int {
	return proto.Size(data)
}

func Decode(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}

func (leafnode *LeafNode) Hash() []byte {
	data, err := proto.Marshal(leafnode)
	if err != nil {
		panic(err)
	}
	return common.Sha256(data)
}

func (innernode *InnerNode) Hash() []byte {
	data, err := proto.Marshal(innernode)
	if err != nil {
		panic(err)
	}
	return common.Sha256(data)
}

func NewErrReceipt(err error) *Receipt {
	berr := err.Error()
	errlog := &ReceiptLog{TyLogErr, []byte(berr)}
	return &Receipt{ExecErr, nil, []*ReceiptLog{errlog}}
}

func CheckAmount(amount int64) bool {
	if amount <= 0 || amount >= MaxCoin {
		return false
	}
	return true
}

func GetEventName(event int) string {
	name, ok := eventName[event]
	if ok {
		return name
	}
	return "unknow-event"
}

func GetSignatureTypeName(signType int) string {
	if name, exist := MapSignType2name[signType]; exist {
		return name
	}

	return "unknow"
}

var ConfigPrefix = "mavl-config-"

func ConfigKey(key string) string {
	return fmt.Sprintf("%s-%s", ConfigPrefix, key)
}

var ManagePrefix = "mavl-"

func ManageKey(key string) string {
	return fmt.Sprintf("%s-%s", ManagePrefix+"manage", key)
}

func ManaeKeyWithHeigh(key string, height int64) string {
	if height >= ForkV13ExecKey {
		return ManageKey(key)
	} else {
		return ConfigKey(key)
	}
}

type ReceiptDataResult struct {
	Ty     int32               `json:"ty"`
	TyName string              `json:"tyname"`
	Logs   []*ReceiptLogResult `json:"logs"`
}

type ReceiptLogResult struct {
	Ty     int32       `json:"ty"`
	TyName string      `json:"tyname"`
	Log    interface{} `json:"log"`
	RawLog string      `json:"rawlog"`
}

func (r *ReceiptData) DecodeReceiptLog() (*ReceiptDataResult, error) {
	result := &ReceiptDataResult{Ty: r.GetTy()}
	switch r.Ty {
	case 0:
		result.TyName = "ExecErr"
	case 1:
		result.TyName = "ExecPack"
	case 2:
		result.TyName = "ExecOk"
	default:
		return nil, ErrLogType
	}

	logs := r.GetLogs()
	for _, l := range logs {
		var lTy string
		var logIns interface{}
		lLog, err := hex.DecodeString(common.ToHex(l.GetLog())[2:])
		if err != nil {
			return nil, err
		}

		logType := LoadLog(int64(l.Ty))
		if logType == nil {
			//tlog.Error("DecodeReceiptLog:", "Faile to decodeLog with type value logtype", l.Ty)
			return nil, ErrLogType
		}

		logIns, err = logType.Decode(lLog)
		lTy = logType.Name()

		result.Logs = append(result.Logs, &ReceiptLogResult{Ty: l.Ty, TyName: lTy, Log: logIns, RawLog: common.ToHex(l.GetLog())})
	}
	return result, nil
}

func (r *ReceiptData) OutputReceiptDetails(logger log.Logger) {
	rds, err := r.DecodeReceiptLog()
	if err == nil {
		logger.Debug("receipt decode", "receipt data", rds)
		for _, rdl := range rds.Logs {
			logger.Debug("receipt log", "log", rdl)
		}
	} else {
		logger.Error("decodelogerr", "err", err)
	}
}

func (t *ReplyGetTotalCoins) IterateRangeByStateHash(key, value []byte) bool {
	//tlog.Debug("ReplyGetTotalCoins.IterateRangeByStateHash", "key", string(key), "value", string(value))
	var acc Account
	err := Decode(value, &acc)
	if err != nil {
		tlog.Error("ReplyGetTotalCoins.IterateRangeByStateHash", "err", err)
		return true
	}
	//tlog.Info("acc:", "value", acc)
	if t.Num >= t.Count {
		t.NextKey = key
		return true
	}
	t.Num++
	t.Amount += acc.Balance
	return false
}

func (action *PrivacyAction) GetInput() *PrivacyInput {
	if action.GetTy() == ActionPrivacy2Privacy && action.GetPrivacy2Privacy() != nil {
		return action.GetPrivacy2Privacy().GetInput()

	} else if action.GetTy() == ActionPrivacy2Public && action.GetPrivacy2Public() != nil {
		return action.GetPrivacy2Public().GetInput()
	}
	return nil
}

func (action *PrivacyAction) GetOutput() *PrivacyOutput {
	if action.GetTy() == ActionPublic2Privacy && action.GetPublic2Privacy() != nil {
		return action.GetPublic2Privacy().GetOutput()
	} else if action.GetTy() == ActionPrivacy2Privacy && action.GetPrivacy2Privacy() != nil {
		return action.GetPrivacy2Privacy().GetOutput()
	} else if action.GetTy() == ActionPrivacy2Public && action.GetPrivacy2Public() != nil {
		return action.GetPrivacy2Public().GetOutput()
	}
	return nil
}

func (action *PrivacyAction) GetActionName() string {
	if action.Ty == ActionPrivacy2Privacy && action.GetPrivacy2Privacy() != nil {
		return "Privacy2Privacy"
	} else if action.Ty == ActionPublic2Privacy && action.GetPublic2Privacy() != nil {
		return "Public2Privacy"
	} else if action.Ty == ActionPrivacy2Public && action.GetPrivacy2Public() != nil {
		return "Privacy2Public"
	}
	return "unknow-privacy"
}

func (action *PrivacyAction) GetTokenName() string {
	if action.GetTy() == ActionPublic2Privacy && action.GetPublic2Privacy() != nil {
		return action.GetPublic2Privacy().GetTokenname()
	} else if action.GetTy() == ActionPrivacy2Privacy && action.GetPrivacy2Privacy() != nil {
		return action.GetPrivacy2Privacy().GetTokenname()
	} else if action.GetTy() == ActionPrivacy2Public && action.GetPrivacy2Public() != nil {
		return action.GetPrivacy2Public().GetTokenname()
	}
	return ""
}

// GetTxTimeInterval 获取交易有效期
func GetTxTimeInterval() time.Duration {
	return time.Second * 120
}

type ParaCrossTx interface {
	IsParaCrossTx() bool
}

func PBToJson(r Message) (string, error) {
	encode := &jsonpb.Marshaler{EmitDefaults: true}
	return encode.MarshalToString(r)
}
