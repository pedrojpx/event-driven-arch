package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (d *EventDispatcher) Register(name string, handler EventHandlerInterface) error {
	if _, ok := d.handlers[name]; ok {
		for _, h := range d.handlers[name] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	d.handlers[name] = append(d.handlers[name], handler)
	return nil
}

func (d *EventDispatcher) Clear() error {
	d.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (d *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := d.handlers[eventName]; ok {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (d *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := d.handlers[event.GetName()]; ok {
		for _, h := range handlers {
			h.Handle(event)
		}
	}
	return nil
}

func (d *EventDispatcher) Unregister(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := d.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				d.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}
	return nil
}
