package main

import (
	"github.com/astaxie/beego/logs"
	"log-agent/tailf"
	"time"
)

func serverRun() (err error) {

	for{
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil{
			logs.Error("send to kafka faild err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg)(err error){
	// 测试将读取到的日志写入项目日志中
	logs.Debug("read msg:%s, topic: %s",msg.Msg, msg.Topic )
	return
}

