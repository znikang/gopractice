package websocketcli

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"webserver/gwebsocket"
)

var (
	StartCmd = &cobra.Command{
		Use:     "websocket",
		Short:   "run websocket server test",
		Example: "webserver websocket",
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

	http.HandleFunc("/ws", gwebsocket.WsHandler)
	fmt.Println("WebSocket server on :8080")
	http.ListenAndServe(":8080", nil)

	return nil
}
