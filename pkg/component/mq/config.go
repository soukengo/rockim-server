package mq

import "rockimserver/pkg/component/mq/kafka"

type Config struct {
	Kafka *kafka.Config
}
