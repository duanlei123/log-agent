package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan chan *TextMsg
}

var(
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) (err error) {
	if len(conf) == 0{
		fmt.Errorf("invalid config for collect, conf: %v", conf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan*TextMsg, chanSize),
	}
	// 遍历数组
	for _, v := range conf{
		obj := &TailObj{
			conf:v,
		}
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		if err != nil {
			err = errTail
			return
		}
		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}
	return
}

//将读取的日志存储到chan
func readFromTail(tailobj *TailObj)  {
	for true {
		msg, ok := <-tailobj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailobj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:   msg.Text,
			Topic: tailobj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}
}
// 从chan获取日志数据
func GetOneLine()(msg *TextMsg){
	msg = <- tailObjMgr.msgChan
	return
}
