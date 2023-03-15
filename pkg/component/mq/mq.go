package mq

import "rockimserver/pkg/component/mq/kafka"

func NewKafkaProducer(cfg *Config) (*kafka.Producer, error) {
	return kafka.NewProducer(cfg.Kafka)
}
