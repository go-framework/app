package event

import (
	"context"
)

type Eventter interface {
	Subscribe(ctx context.Context, event string, callback func(context.Context, ...interface{}) error)
	Publish(ctx context.Context, event string, args ...interface{}) error
	Unsubscribe(event string, callback ...func(context.Context, ...interface{}) error)
}
