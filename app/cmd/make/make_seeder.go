package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example:  make seeder user",
	Args:  cobra.ExactArgs(1),
	Run:   runMakeSeeder,
}

func runMakeSeeder(cmd *cobra.Command, args []string) {

	//格式化模型名称,返回一个 model 对象
	model := makeModelFromString(args[0])

	//拼接目标文件路径
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.PackageName)

	//基于模版创建文件
	createFileFromStub(filePath, "seeder", model)
}
