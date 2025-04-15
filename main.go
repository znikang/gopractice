package main

import (
	"webserver/cmd"
	"webserver/common"
)

func main() {

	common.InitLogger(common.Config{
		EnableFile:  true,
		LogFilePath: "app.log",
		UseJSON:     true,
		LogLevel:    common.Log().Level, // or logrus.InfoLevel
	})

	cmd.Execute()

	defer common.CloseLogFile()

	common.Log().Info("app started")
	common.Log().Warn("this is a warning")
	common.Log().Error("this is an error")

}
