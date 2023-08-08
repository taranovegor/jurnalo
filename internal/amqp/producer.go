package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/taranovegor/com.jurnalo/internal/config"
	"reflect"
)

type Producer struct {
	producers config.AmqpProducers
	channel   *amqp.Channel
}

func NewProducer(
	producers config.AmqpProducers,
	channel *amqp.Channel,
) *Producer {
	return &Producer{
		producers: producers,
		channel:   channel,
	}
}

func (prod Producer) Publish(msg interface{}) {
	typeOfMsg := reflect.TypeOf(msg)
	for msgType, queues := range prod.producers {
		if typeOfMsg != reflect.TypeOf(msgType) {
			continue
		}

		serialized, _ := json.Marshal(msg)
		for _, queue := range queues {
			prod.channel.QueueDeclare(queue, false, false, false, false, nil)
			prod.channel.Publish(
				"",
				queue,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        serialized,
				},
			)
		}
	}
}
