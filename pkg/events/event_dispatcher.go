package events

import (
	"errors"
	"sync"
)

var errHandlerAlreadyRegister = errors.New("Handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if ed.Has(eventName, handler) {
		return errHandlerAlreadyRegister
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil

}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil

}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handles, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handles {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = map[string][]EventHandlerInterface{}
	return nil
}
