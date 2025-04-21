package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
)

type Consumer[T any] struct {
	consumer sarama.ConsumerGroup
	topic    string
}

func NewConsumer[T any](brokers []string, groupID string, topic string) (*Consumer[T], error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &Consumer[T]{
		consumer: consumerGroup,
		topic:    topic,
	}, nil
}

func (c *Consumer[T]) Consume(ctx context.Context, handler func(event T)) error {
	// Create handler instance with proper interface implementation
	cgHandler := &consumerGroupHandler[T]{
		callback: handler,
		ready:    make(chan bool),
	}

	// Start consuming in a loop to handle rebalances
	go func() {
		for {
			if err := c.consumer.Consume(ctx, []string{c.topic}, cgHandler); err != nil {
				return
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	// Wait until the consumer has been set up
	<-cgHandler.ready
	return nil
}

type consumerGroupHandler[T any] struct {
	callback func(T)
	ready    chan bool
}

func (h *consumerGroupHandler[T]) Setup(sess sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *consumerGroupHandler[T]) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler[T]) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var event T
		if err := json.Unmarshal(message.Value, &event); err != nil {
			continue
		}
		h.callback(event)
		sess.MarkMessage(message, "")
	}
	return nil
}

func (c *Consumer[T]) Close() error {
	return c.consumer.Close()
}
