package main

import (
	"webserver/cmd"
	"webserver/common"
)

func main() {

	cmd.Execute()
	defer common.GrpcCli.CloseRpc()
}
