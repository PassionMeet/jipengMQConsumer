package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	nsqConf := nsq.NewConfig()
	var err error
	consumer, err := nsq.NewConsumer("user_geo", "match_man", nsqConf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(&userGeo_matchMan_MessageHandler{}, 10)
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		panic(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
