package producers

var UserProducer *kafka.Producer

func InitUserProducer() error {
	var err error
	UserProducer, err = kafka.NewProducer(
		cfg.Brokers,
		cfg.Topics.UserEvents,
	)
	return err
}

func CloseUserProducer() {
	if UserProducer != nil {
		UserProducer.Close()
	}
}
