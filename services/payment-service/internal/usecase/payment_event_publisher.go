package usecase

type PaymentEventPublisher interface {
	PublishMessage(topic, message string) error
}
