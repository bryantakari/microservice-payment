package order

import "context"

type EventPublisher interface {
	PublishOrderCreated(ctx context.Context, event OrderCreatedEvent) error
}

type OrderCreatedEvent struct {
	OrderID string
	UserID  string
	Amount  int64
}
