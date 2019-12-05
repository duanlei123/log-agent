package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"log-agent/tailf"
)

var (
	appConfig *Config
)

type Config struct {
	LogLevel    string              //日志级别
	LogPath     string              // 日志路径
	CollectConf []tailf.CollectConf //所收集的日志
}

func loadCollectConf(conf config.Configer) (err error) {
	var collectConf tailf.CollectConf
	collectConf.LogPath = conf.String("collect::log_path")
	if len(collectConf.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}
	collectConf.Topic = conf.String("collect::topic")
	if len(collectConf.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}
	appConfig.CollectConf = append(appConfig.CollectConf, collectConf)
	return
}

func loadConf(confType, filename string) error {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err :", err)
		return err
	}
	appConfig = &Config{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		// 设置兜底默认值
		appConfig.LogLevel = "dubug"
	}
	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		// 设置兜底默认值
		appConfig.LogPath = "./logs"
	}
	// 加载收集相关的配置
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Println("load Collect config failed, err :", err)
		return err
	}
	return nil
}
