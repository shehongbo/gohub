package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个model 对象
	model := makeModelFromString(args[0])

	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll("app/policies", os.ModePerm)

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	//基于模版创建文件
	createFileFromStub(filePath, "policy", model)
}
