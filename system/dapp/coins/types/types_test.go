package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.33.cn/chain33/chain33/types"
)

func TestTypeReflact(t *testing.T) {
	ty := NewType()
	assert.NotNil(t, ty)

	//创建一个json字符串
	data, err := types.PBToJson(&types.AssetsTransfer{Amount: 10})
	assert.Nil(t, err)
	raw := json.RawMessage(data)
	tx, err := ty.CreateTx("Transfer", raw)
	assert.Nil(t, err)

	name, val, err := ty.DecodePayloadValue(tx)
	assert.Nil(t, err)
	assert.Equal(t, name, "Transfer")
	assert.Equal(t, val.Interface().(*types.AssetsTransfer).GetAmount(), int64(10))
}
