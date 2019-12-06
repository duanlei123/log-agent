package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var(
	client sarama.SyncProducer
)

func InitKafkaClient(addr string)(err error){
	//配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应,即数据成功发送到kafka后返回的响应信息
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true

	//使用配置,新建一个异步生产者
	//sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed , err :",err)
		return
	}
	logs.Debug("init kafka success")
	return
}

func SendToKafka(data, topic string)(err error){
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed,err:%v data:%v topic:%v", err, data, topic)
		return
	}
	logs.Debug("send success, topic:%v pid: %v offset: %v",topic, pid, offset)
	return
}
