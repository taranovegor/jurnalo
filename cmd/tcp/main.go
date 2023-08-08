package main

import (
	"github.com/taranovegor/com.jurnalo/cmd"
	"github.com/taranovegor/com.jurnalo/internal/amqp"
	"github.com/taranovegor/com.jurnalo/internal/container"
	kontrakto "github.com/taranovegor/com.kontrakto"
)

func main() {
	sc := cmd.Init(container.ScopeTcp)

	producer := sc.Get(container.AmqpProducer).(*amqp.Producer)
	producer.Publish(kontrakto.SendNotification{Message: "!!!"})

	//handler := sc.Get(container.HandlerTcp).(tcp.Handler)
	//handler.Handle()
}
