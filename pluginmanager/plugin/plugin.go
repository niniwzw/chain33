package plugin

import (
	"gitlab.33.cn/chain33/chain33/types"

	"github.com/spf13/cobra"
)

//
type Plugin interface {
	// 获取整个插件的包名，用以计算唯一值、做前缀等
	GetPackageName() string
	// 获取插件中执行器名
	GetExecutorName() string
	// 初始化执行器时会调用该接口
	InitExecutor()
	DecodeTx(tx *types.Transaction) interface{}
	AddCustomCommand(rootCmd *cobra.Command)
}
