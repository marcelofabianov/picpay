package port

import "context"

type BrokerMessage interface {
	Publish(ctx context.Context, exchange string, key string, message []byte) error
	Consumer(ctx context.Context, exchange string, key string, handler func(message []byte) error) error
}
