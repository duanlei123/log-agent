package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log-agent/tailf"
)

func main() {
	//初始化配置文件
	filename := "src/log-agent/conf/logAgent.ini"
	err := loadConf("ini",filename)
	if err != nil{
		fmt.Println("load conf failed err :", err)
		panic("load conf failed")
		return
	}
	//初始化项目日志
	err = initLogger()
	if err != nil{
		fmt.Println("load logger failed err :", err)
		panic("load logger failed")
		return
	}

	err = tailf.InitTail(appConfig.CollectConf, appConfig.ChanSize)
	if err != nil{
		logs.Error("init Tail failed err :", err)
		return
	}
	logs.Debug("initAll success")

	err = serverRun()
	if err != nil{
		logs.Error("serverRun failed err:", err)
		return
	}
	logs.Info("program exited")
}
