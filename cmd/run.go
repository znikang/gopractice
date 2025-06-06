package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"webserver/cmd/authserver"
	croncli "webserver/cmd/cron"
	"webserver/cmd/grpcclient"
	"webserver/cmd/grpcserver"
	"webserver/cmd/nacos"
	ormserver "webserver/cmd/orm"
	"webserver/cmd/server"
	"webserver/cmd/version"
	websocketcli "webserver/cmd/websocket"
)

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	usageStr1 := `也可以参考  的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var rootCmd = &cobra.Command{
	Use:          "webserver",
	Short:        "",
	SilenceUsage: true,
	Long:         `webserver `,
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
	rootCmd.AddCommand(grpcserver.StartCmd)
	rootCmd.AddCommand(grpcclient.StartCmd)
	rootCmd.AddCommand(croncli.StartCmd)
	rootCmd.AddCommand(websocketcli.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
