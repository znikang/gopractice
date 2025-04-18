package main

import (
	"github.com/sirupsen/logrus"
	"webserver/cmd"
	"webserver/common"
)

func main() {

	cf := common.Config{
		EnableFile:  true,
		LogFilePath: "app.log",
		UseJSON:     true,
		LogLevel:    logrus.DebugLevel, // or logrus.InfoLevel
	}
	common.InitLogger(cf)

	cmd.Execute()

	defer common.CloseLogFile()

	//common.Log().Info("app started")
	//common.Log().Warn("this is a warning")
	//common.Log().Error("this is an error")

}
