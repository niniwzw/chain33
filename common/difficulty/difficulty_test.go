package difficulty

import (
	"math/big"
	"strings"
	"testing"

	"gitlab.33.cn/chain33/chain33/common"
)

func TestCompact(t *testing.T) {
	hashhex := "0x0000" + strings.Repeat("F", 60)
	b, err := common.FromHex(hashhex)
	if err != nil {
		t.Log(err)
		return
	}
	num := big.NewInt(0).SetBytes(b)
	bits := BigToCompact(num)
	t.Logf("%x", bits)
}
