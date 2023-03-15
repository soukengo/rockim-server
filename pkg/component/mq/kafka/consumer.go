package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"rockimserver/pkg/log"
	"rockimserver/pkg/util/runtimes"
	"strings"
	"sync"
)

const (
	defaultWorkers = 10
)

type Consumer struct {
	consumer *cluster.Consumer
	config   *ConsumerConfig
	handler  HandlerFunc
	queue    chan *sarama.ConsumerMessage // 数据队列
	consumed sync.Once
	closed   sync.Once
	done     chan struct{}
}

type HandlerFunc func(topic string, body []byte)

func NewConsumer(config *ConsumerConfig, topics []string, handler HandlerFunc) (*Consumer, error) {
	kafkaConfig := cluster.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Group.Return.Notifications = true
	for i := 0; i < len(topics); i++ {
		topics[i] = config.Kafka.TopicPrefix + topics[i]
	}
	if config.Workers <= 0 {
		config.Workers = defaultWorkers
	}
	consumer, err := cluster.NewConsumer(config.Kafka.Brokers, config.Group, topics, kafkaConfig)
	if err != nil {
		return nil, err
	}
	c := &Consumer{
		config:   config,
		consumer: consumer,
		handler:  handler,
		queue:    make(chan *sarama.ConsumerMessage, config.Workers),
		done:     make(chan struct{}),
	}
	return c, nil
}

// Consume 消费
func (s *Consumer) Consume() {
	s.consumed.Do(func() {
		go runtimes.TryCatch(func() {
			s.receive()
		})
		go runtimes.TryCatch(func() {
			s.dispatch()
		})
	})
}

// dispatch 分发
func (s *Consumer) dispatch() {
	for {
		select {
		case <-s.done:
			return
		case msg := <-s.queue:
			topic := msg.Topic
			// 移除配置的topic前缀
			topic = strings.TrimPrefix(topic, s.config.Kafka.TopicPrefix)
			s.handler(topic, msg.Value)
		}
	}
}

// receive 接收
func (s *Consumer) receive() {
	for {
		select {
		case <-s.done:
			return
		case err := <-s.consumer.Errors():
			log.Errorf("consumer error(%v)", err)
		case n := <-s.consumer.Notifications():
			log.Infof("consumer rebalanced(%v)", n)
		case msg := <-s.consumer.Messages():
			s.consumer.MarkOffset(msg, "")
			s.queue <- msg
		}
	}
}

func (s *Consumer) Close() error {
	s.closed.Do(func() {
		close(s.done)
		close(s.queue)
	})
	return nil
}
