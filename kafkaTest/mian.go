package kafkaTest

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//kafka 示例代码
func main() {
	//配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应,即数据成功发送到kafka后返回的响应信息
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	fmt.Println("start make producer")

	//使用配置,新建一个异步生产者
	//sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	client, err := sarama.NewSyncProducer([]string{"192.168.18.131:9092"}, config)
	if err != nil {
		fmt.Println("product close, err :", err)
		return
	}

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	// 创建消息
	defer client.Close()
	fmt.Printf("pid: %v offset: %v\n", pid, offset)
}
