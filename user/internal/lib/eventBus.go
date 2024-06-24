package lib

import (
	eventbus "github.com/dtomasi/go-event-bus/v3"
)

type EventInterface interface {
	GetEventName() string
}

type EventHandler interface {
	Subscribed() EventInterface
	Execute(interface{})
}

type EventBusInterface interface {
	Subscribe(EventHandler)
	Dispatch(EventInterface)
}

type EventBus struct {
	bus *eventbus.EventBus
	//events []EventHandler
}

func (b EventBus) Subscribe(handler EventHandler) {
	b.bus.SubscribeCallback(
		handler.Subscribed().GetEventName(),
		func(topic string, data interface{}) {
			handler.Execute(data)
		},
	)
}

func (b EventBus) Dispatch(input EventInterface) {
	b.bus.Publish(input.GetEventName(), input)
}

var EventBusInstance *EventBus

func init() {
	//fmt.Println("Event bus init")
	EventBusInstance = &EventBus{
		bus: eventbus.NewEventBus(),
	}
}
