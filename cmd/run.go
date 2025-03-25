package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"webserver/cmd/authserver"
	"webserver/cmd/nacos"
	ormserver "webserver/cmd/orm"
	"webserver/cmd/server"
	"webserver/cmd/version"
)

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	usageStr1 := `也可以参考  的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var rootCmd = &cobra.Command{
	Use:          "mysql",
	Short:        "mysql",
	SilenceUsage: true,
	Long:         `mysql`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("請輸入參數")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {

	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(nacos.StartCmd)
	rootCmd.AddCommand(ormserver.StartCmd)
	rootCmd.AddCommand(authserver.StartCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
