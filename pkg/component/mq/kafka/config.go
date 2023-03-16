package kafka

import "rockimserver/pkg/component/mq"

type ConsumerConfig struct {
	Group   string
	Topics  []string
	Workers int32
	Kafka   *mq.Kafka
}
