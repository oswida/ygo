package basic

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	defaultTimeout         = time.Millisecond * 500
	defaultChannelCapacity = 100
)

type Subscriber[T any] interface {
	Close()
	Events() <-chan T
	Emit(ctx context.Context, event T) error
}

type Observable[T any] interface {
	Emit(ctx context.Context, event T)
	Subscribe() Subscriber[T]
}

type observable[T any] struct {
	observers SyncMap[uuid.UUID, *subscriber[T]]
}

func NewObservable[T any]() Observable[T] {
	return &observable[T]{
		observers: SyncMap[uuid.UUID, *subscriber[T]]{},
	}
}

func (s *observable[T]) Emit(ctx context.Context, event T) {
	subscribers := s.observers.Values()
	var timedOutSubscribers []*subscriber[T]
	for i := range subscribers {
		if err := subscribers[i].Emit(ctx, event); err != nil {
			timedOutSubscribers = append(timedOutSubscribers, subscribers[i])
			continue
		}
	}

	for i := range timedOutSubscribers {
		timedOutSubscribers[i].Close()
	}
}

func (s *observable[T]) Subscribe() Subscriber[T] {
	subscription := &subscriber[T]{
		id:         uuid.New(),
		events:     make(chan T, defaultChannelCapacity),
		observable: s,
	}
	s.observers.Set(subscription.id, subscription)
	return subscription
}

type subscriber[T any] struct {
	id         uuid.UUID
	events     chan T
	observable *observable[T]
}

func (s *subscriber[T]) Close() {
	s.observable.observers.Delete(s.id)
}

func (s *subscriber[T]) Events() <-chan T {
	return s.events
}

func (s *subscriber[T]) Emit(ctx context.Context, data T) error {
	select {
	case s.events <- data:
	case <-ctx.Done():
		return nil
	case <-time.After(defaultTimeout):
		return os.ErrDeadlineExceeded
	}

	return nil
}
