package croncli

import (
	"github.com/spf13/cobra"
	"webserver/common"
)

var (
	StartCmd = &cobra.Command{
		Use:     "cron",
		Short:   "run cronjob client test",
		Example: "webserver cron",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {

}

func initTools() {

}

func run() error {

	common.Log().Println("dasda")
	return nil
}
