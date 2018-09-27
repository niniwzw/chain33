package executor

import (
	pty "gitlab.33.cn/chain33/chain33/plugin/dapp/hashlock/types"
	"gitlab.33.cn/chain33/chain33/types"
)

func (h *Hashlock) ExecDelLocal_Hlock(transfer *pty.HashlockLock, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return nil, nil
}

func (h *Hashlock) ExecDelLocal_Hsend(transfer *pty.HashlockSend, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return nil, nil
}

func (h *Hashlock) ExecDelLocal_Hunlock(transfer *pty.HashlockUnlock, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return nil, nil
}
