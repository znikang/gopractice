package kafkahook

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaHook struct {
	Writer *kafka.Writer
	Topic  string
}

func NewKafkaHook(brokerAddr string, topic string) *KafkaHook {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaHook{
		Writer: writer,
		Topic:  topic,
	}
}

func (hook *KafkaHook) Fire(entry *logrus.Entry) error {
	payload, err := json.Marshal(entry.Data)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Value: payload,
	}
	return hook.Writer.WriteMessages(context.Background(), msg)
}

func (hook *KafkaHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
