package events

import "time"

type EventInterface interface {
	GetName() string
	GetDate() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handler(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatcher(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
