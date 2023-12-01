package events

import "errors"

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

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errHandlerAlreadyRegister
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil

}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = map[string][]EventHandlerInterface{}
	return nil
}
