package subpub

import (
	"context"
	"sync"
)

type MessageHandler func(msg interface{})

type Subscription struct {
	cancel func()
}

func (s *Subscription) Unsubscribe() {
	if s.cancel != nil {
		s.cancel()
	}
}

type SubPub interface {
	Subscribe(subject string, cb MessageHandler) (Subscription, error)
	Publish(subject string, msg interface{}) error
	Close(ctx context.Context) error
}

type subPub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan interface{}
	closed      bool
}

func NewSubPub() SubPub {
	return &subPub{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (sp *subPub) Subscribe(subject string, cb MessageHandler) (Subscription, error) {
	if sp.closed {
		return Subscription{}, context.Canceled
	}

	sp.mu.Lock()
	defer sp.mu.Unlock()

	// Создание буферизированного канала для сообщений
	msgChan := make(chan interface{}, 100)

	// Добавление канала в список подписчиков
	sp.subscribers[subject] = append(sp.subscribers[subject], msgChan)

	// Запуск горутины для обработки сообщений
	go func() {
		for msg := range msgChan {
			cb(msg)
		}
	}()

	// Создание функции отписки
	cancel := func() {
		sp.mu.Lock()
		defer sp.mu.Unlock()

		// Находим и удаляем канал из списка
		channels := sp.subscribers[subject]
		for i, ch := range channels {
			if ch == msgChan {
				close(ch)
				sp.subscribers[subject] = append(channels[:i], channels[i+1:]...)
				break
			}
		}

		// Если подписчиков больше нет, то удаляем subject
		if len(sp.subscribers[subject]) == 0 {
			delete(sp.subscribers, subject)
		}
	}

	return Subscription{cancel: cancel}, nil
}

func (sp *subPub) Publish(subject string, msg interface{}) error {
	if sp.closed {
		return context.Canceled
	}

	sp.mu.RLock()
	channels := sp.subscribers[subject]
	sp.mu.RUnlock()

	// Отправляем сообщение всем подписчикам
	for _, ch := range channels {
		select {
		case ch <- msg:
			// Сообщение отправлено
		default:
			// Канал заполнен, пропускаем
		}
	}

	return nil
}

func (sp *subPub) Close(ctx context.Context) error {
	sp.mu.Lock()
	defer sp.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if !sp.closed {
			sp.closed = true
			// Закрываем все каналы
			for _, channels := range sp.subscribers {
				for _, ch := range channels {
					close(ch)
				}
			}
			sp.subscribers = nil
		}
		return nil
	}
}
