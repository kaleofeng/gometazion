package eventbus

import (
	"reflect"
	"sync"
)

type EventBus struct {
	subs  map[string][]Subscriber
	mutex sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subs: make(map[string][]Subscriber),
	}
}

func (e *EventBus) Register(evt string, sub Subscriber) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.subs[evt] = append(e.subs[evt], sub)
}

func (e *EventBus) Unregister(evt string, sub Subscriber) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	subs := e.subs[evt]
	for i := 0; i < len(subs); {
		if reflect.ValueOf(subs[i]) == reflect.ValueOf(sub) {
			subs = append(subs[:i], subs[i+1:]...)
		} else {
			i++
		}
	}
	e.subs[evt] = subs
}

func (e *EventBus) Emit(evt string, data interface{}) {
	for _, sub := range e.subs[evt] {
		sub.OnEvent(evt, data)
	}
}
