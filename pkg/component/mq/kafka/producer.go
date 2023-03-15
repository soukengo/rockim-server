package kafka

import (
	kafka "github.com/Shopify/sarama"
)

type Producer struct {
	brokers     []string
	topicPrefix string
	producer    kafka.SyncProducer
}

func NewProducer(c *Config) (*Producer, error) {
	ins := &Producer{brokers: c.Brokers, topicPrefix: c.TopicPrefix}
	kc := kafka.NewConfig()
	kc.Producer.RequiredAcks = kafka.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                  // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafka.NewSyncProducer(ins.brokers, kc)
	if err != nil {
		return nil, err
	}
	ins.producer = pub
	return ins, nil
}

func (c *Producer) Close() error {
	return c.producer.Close()
}

func (c *Producer) SendMessage(topic string, key string, bytes []byte) (partition int32, offset int64, err error) {
	m := &kafka.ProducerMessage{
		Key:   kafka.StringEncoder(key),
		Topic: c.topicPrefix + topic,
		Value: kafka.ByteEncoder(bytes),
	}
	return c.producer.SendMessage(m)
}
