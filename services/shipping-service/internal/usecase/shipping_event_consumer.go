package usecase

type ShippingEventConsumer interface {
	StartConsuming(workerCount int)
}
