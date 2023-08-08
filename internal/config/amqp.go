package config

import kontrakto "github.com/taranovegor/com.kontrakto"

type AmqpProducers map[interface{}][]string

type AmqpConfig struct {
	Producers AmqpProducers
}

func GetAmqpConfig() AmqpConfig {
	return AmqpConfig{
		Producers: map[interface{}][]string{
			kontrakto.SendNotification{}: {kontrakto.QueueNotification},
		},
	}
}
