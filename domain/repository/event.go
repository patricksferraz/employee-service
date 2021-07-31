package repository

import (
	"context"
)

type EventRepositoryInterface interface {
	Publish(ctx context.Context, msg, topic, key string) error
}
