package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/edorguez/payment-reminder/pkg/kafka/events"
)

type Consumer struct {
	consumer sarama.ConsumerGroup
	topic    string
}

func NewConsumer(brokers []string, groupID string, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumerGroup,
		topic:    topic,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context, handler func(event any)) error {
	// Create handler instance with proper interface implementation
	cgHandler := &consumerGroupHandler{
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

type consumerGroupHandler struct {
	callback func(any)
	ready    chan bool
}

func (h *consumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *consumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var event events.UserEvent
		if err := json.Unmarshal(message.Value, &event); err != nil {
			continue
		}
		h.callback(event)
		sess.MarkMessage(message, "")
	}
	return nil
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
