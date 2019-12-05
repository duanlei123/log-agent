package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

//git上log日志组件
//https://github.com/hpcloud/tail/blob/master/tail.go
func main() {
	filename := "C:\\Users\\xxxxxx\\Downloads\\access.log"

	config := tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
