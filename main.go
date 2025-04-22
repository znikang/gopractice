package main

import (
	"github.com/sirupsen/logrus"
	"webserver/cmd"
	"webserver/common"
	kafkahook "webserver/kafka"
)

func main() {

	cf := common.Config{
		EnableFile:  true,
		LogFilePath: "app.log",
		UseJSON:     true,
		LogLevel:    logrus.DebugLevel, // or logrus.InfoLevel
	}
	common.InitLogger(cf)

	kafkaHook := kafkahook.NewKafkaHook("192.168.0.146:9092", "logrus-topic")
	common.Log().AddHook(kafkaHook)

	cmd.Execute()

	defer common.CloseLogFile()

	//common.Log().Info("app started")
	//common.Log().Warn("this is a warning")
	//common.Log().Error("this is an error")

}
