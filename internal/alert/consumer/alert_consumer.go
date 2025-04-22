package consumer

import (
	"context"

	"github.com/edorguez/payment-reminder/internal/alert/repository"
	"github.com/edorguez/payment-reminder/pkg/constants"
	"github.com/edorguez/payment-reminder/pkg/kafka"
	"github.com/edorguez/payment-reminder/pkg/kafka/events"
)

type AlertConsumer struct {
	userCacheRepo repository.UserCacheRepository
	kafkaTopic    string
	groupID       string
}

func NewAlertConsumer(userCacheRepo repository.UserCacheRepository) *AlertConsumer {
	return &AlertConsumer{
		userCacheRepo: userCacheRepo,
		kafkaTopic:    "users",
		groupID:       "alert-service-group",
	}
}

func (c *AlertConsumer) Start(brokers []string) error {
	consumer, err := kafka.NewConsumer[events.UserEvent](
		brokers,
		c.groupID,
		c.kafkaTopic,
	)
	if err != nil {
		return err
	}

	go c.consumeMessages(consumer)
	return nil
}

func (c *AlertConsumer) consumeMessages(consumer *kafka.Consumer[events.UserEvent]) {
	err := consumer.Consume(context.Background(), func(event events.UserEvent) {

		switch event.EventType {
		case constants.UserCreatedEvent:
			userCache := repository.UserCache{
				ID:    event.UserID,
				Email: event.Email,
			}
			c.userCacheRepo.Create(context.Background(), userCache)
		case constants.UserUpdatedEvent:
			userCache := repository.UserCache{
				ID:    event.UserID,
				Email: event.Email,
			}
			c.userCacheRepo.Update(context.Background(), userCache)
		case constants.UserDeletedEvent:
			c.userCacheRepo.Delete(context.Background(), event.UserID)
		}
	})

	if err != nil {
		// Handle error properly (don't panic in production)
		panic(err)
	}
}
