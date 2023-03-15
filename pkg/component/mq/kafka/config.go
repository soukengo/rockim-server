package kafka

type Config struct {
	Brokers     []string
	TopicPrefix string
}

type ConsumerConfig struct {
	Group   string
	Workers int32
	Kafka   *Config
}
