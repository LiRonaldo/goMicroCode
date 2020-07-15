package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/mqtt"
	_ "github.com/micro/go-plugins/broker/mqtt"
	"goMicroCode/message"
	"log"
)

func main() {

	//实例化
	service := micro.NewService(
		micro.Name("go.micro.srv"),
		micro.Version("v1.0.0"),
		micro.Broker(mqtt.NewBroker()),
	)

	//初始化
	service.Init()

	//订阅事件
	pubSub := service.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatal(" broker connect failed , error is : %v\n", err)
	}

	_, err := pubSub.Subscribe("go.micro.srv.message", func(event broker.Event) error {
		var req *message.StudentRequest
		fmt.Println(string(event.Message().Body))

		if err := json.Unmarshal(event.Message().Body, &req); err != nil {
			return err
		}
		fmt.Println(" 接收到信息：", req)
		//去执行其他操作
		return nil
	})

	if err != nil {
		log.Printf("sub error: %v\n", err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
