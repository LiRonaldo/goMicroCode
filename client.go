package main

import (
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/mqtt"
	_ "github.com/micro/go-plugins/broker/mqtt"
	"goMicroCode/message"
	"log"
)

func main() {

	service := micro.NewService(
		micro.Name("student.client"),
		micro.Broker(mqtt.NewBroker()),
	)

	service.Init()

	brok := service.Server().Options().Broker
	if err := brok.Connect(); err != nil {
		log.Fatal(" broker connection failed, error : ", err.Error())
	}

	student := &message.Student{Name: "李玉祥", Classes: "软件工程专业", Grade: 80, Phone: "12345678901"}
	msgBody, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err.Error())
	}

	msg := &broker.Message{
		Header: map[string]string{
			"name": student.Name,
		},
		Body: msgBody,
	}

	err = brok.Publish("go.micro.srv.message", msg)
	if err != nil {
		log.Fatal(" 消息发布失败：%s\n", err.Error())
	} else {
		log.Print("消息发布成功")
	}

	defer brok.Disconnect()
}
