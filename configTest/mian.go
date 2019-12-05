package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "D:\\go_code\\project01\\src\\log-agent\\configTest\\test.ini")
	if err != nil {
		fmt.Println("new config failed, err :", err)
		return
	}
	port, err := conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("read server::port failed, err:", err)
		return
	}
	fmt.Println("port:", port)

	log_level, err := conf.Int("logs::log_level")
	if err != nil {
		fmt.Println("read logs::log_level failed, err:", err)
		return
	}
	fmt.Println("log_level:", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log_path:", log_path)

}
