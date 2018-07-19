package paracross

import (
	"gitlab.33.cn/chain33/chain33/types"
	"fmt"
)

var (
	title string
	titleHeight string
)

func setPrefix() {
	title = "mavl-" + types.ExecName("paracross") + "title-"
	titleHeight = "mavl-" + types.ExecName("paracross") + "titleHeight-"
}

func calcTitleKey(title string) []byte {
	return []byte(fmt.Sprintf(title+"%s", title))
}

func calcTitleHeightKey(title string, height int64) []byte {
	return []byte(fmt.Sprintf(titleHeight+"%s-%012d", title, height))
}


