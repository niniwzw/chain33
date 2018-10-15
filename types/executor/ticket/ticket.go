package ticket

import (
	"encoding/json"

	//log "github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/types"
)

var nameX string

var (
	actionName = map[string]int32{
		"Genesis": types.TicketActionGenesis,
		"Topen":   types.TicketActionOpen,
		"Tbind":   types.TicketActionBind,
		"Tclose":  types.TicketActionClose,
		"Miner":   types.TicketActionMiner,
	}
)

//var tlog = log.New("module", name)

func Init() {
	nameX = types.ExecName("ticket")
	// init executor type
	types.RegistorExecutor("ticket", NewType())

	// init log
	types.RegistorLog(types.TyLogNewTicket, &TicketNewLog{})
	types.RegistorLog(types.TyLogCloseTicket, &TicketCloseLog{})
	types.RegistorLog(types.TyLogMinerTicket, &TicketMinerLog{})
	types.RegistorLog(types.TyLogTicketBind, &TicketBindLog{})
}

type TicketType struct {
	types.ExecTypeBase
}

func NewType() *TicketType {
	c := &TicketType{}
	c.SetChild(c)
	return c
}

func (at *TicketType) GetPayload() types.Message {
	return &types.TicketAction{}
}

func (ticket TicketType) ActionName(tx *types.Transaction) string {
	var action types.TicketAction
	err := types.Decode(tx.Payload, &action)
	if err != nil {
		return "unknown-err"
	}
	if action.Ty == types.TicketActionGenesis && action.GetGenesis() != nil {
		return "genesis"
	} else if action.Ty == types.TicketActionOpen && action.GetTopen() != nil {
		return "open"
	} else if action.Ty == types.TicketActionClose && action.GetTclose() != nil {
		return "close"
	} else if action.Ty == types.TicketActionMiner && action.GetMiner() != nil {
		return types.MinerAction
	} else if action.Ty == types.TicketActionBind && action.GetTbind() != nil {
		return "bindminer"
	}
	return "unknown"
}

func (ticket TicketType) DecodePayload(tx *types.Transaction) (interface{}, error) {
	var action types.TicketAction
	err := types.Decode(tx.Payload, &action)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (ticket TicketType) Amount(tx *types.Transaction) (int64, error) {
	var action types.TicketAction
	err := types.Decode(tx.GetPayload(), &action)
	if err != nil {
		return 0, types.ErrDecode
	}
	if action.Ty == types.TicketActionMiner && action.GetMiner() != nil {
		ticketMiner := action.GetMiner()
		return ticketMiner.Reward, nil
	}
	return 0, nil
}

// TODO 暂时不修改实现， 先完成结构的重构
func (ticket TicketType) CreateTx(action string, message json.RawMessage) (*types.Transaction, error) {
	var tx *types.Transaction
	return tx, nil
}

func (ticket TicketType) GetName() string {
	return "ticket"
}

func (ticket *TicketType) GetLogMap() map[int64]*types.LogInfo {
	return nil
}

func (ticket *TicketType) GetTypeMap() map[string]int32 {
	return actionName
}

type TicketNewLog struct {
}

func (l TicketNewLog) Name() string {
	return "LogNewTicket"
}

func (l TicketNewLog) Decode(msg []byte) (interface{}, error) {
	var logTmp types.ReceiptTicket
	err := types.Decode(msg, &logTmp)
	if err != nil {
		return nil, err
	}
	return logTmp, err
}

type TicketCloseLog struct {
}

func (l TicketCloseLog) Name() string {
	return "LogCloseTicket"
}

func (l TicketCloseLog) Decode(msg []byte) (interface{}, error) {
	var logTmp types.ReceiptTicket
	err := types.Decode(msg, &logTmp)
	if err != nil {
		return nil, err
	}
	return logTmp, err
}

type TicketMinerLog struct {
}

func (l TicketMinerLog) Name() string {
	return "LogMinerTicket"
}

func (l TicketMinerLog) Decode(msg []byte) (interface{}, error) {
	var logTmp types.ReceiptTicket
	err := types.Decode(msg, &logTmp)
	if err != nil {
		return nil, err
	}
	return logTmp, err
}

type TicketBindLog struct {
}

func (l TicketBindLog) Name() string {
	return "LogTicketBind"
}

func (l TicketBindLog) Decode(msg []byte) (interface{}, error) {
	var logTmp types.ReceiptTicketBind
	err := types.Decode(msg, &logTmp)
	if err != nil {
		return nil, err
	}
	return logTmp, err
}
