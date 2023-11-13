package main

import (
	"dingchatgpt/config"
	"dingchatgpt/db"
	"dingchatgpt/ding"
	"flag"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitFlat()

	flag.Parse()
	//db.GptDBModel().Init()
	db.Connect()

	ding.StartDingTalk()

	signChan := make(chan os.Signal, 1)

	signal.Notify(signChan, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt)

	<-signChan

	logger.GetLogger().Infof("app closed")
}

