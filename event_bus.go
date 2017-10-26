package pong

import (
	"github.com/chuckpreslar/emission"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"fmt"
	"errors"
	"encoding/json"
)

const DefaultConsulUri = "http://localhost:8500"

type EventBus struct {
	*emission.Emitter

	consulUri string
	consulClient *api.Client
	consulWatch  *watch.Plan

	seen uint64
}

func NewEventBus(consulUri string) *EventBus {
	eb := &EventBus{Emitter: emission.NewEmitter()}
	if consulUri != "" {
		eb.consulUri = consulUri
	} else {
		eb.consulUri = DefaultConsulUri
	}

	return eb
}

func (e *EventBus) Start() error {
	var err error
	e.consulClient, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	// setup raw event handler
	e.On("rawevent", e.handle)

    // setup watch for pong events
	watchParams := make(map[string]interface{})
	watchParams["type"] = "event"
	watchParams["name"] = "pong"

	e.consulWatch, err = watch.Parse(watchParams)
	if err != nil {
		return err
	}

	// Set handler
	e.consulWatch.Handler = func(idx uint64, data interface{}) {
		events := data.([]*api.UserEvent)

		for _, event := range events {
			if event.LTime > e.seen {
				e.Emit("rawevent", event)
				e.seen = event.LTime
			}
		}
	}

	if err := e.consulWatch.Run(e.consulUri); err != nil {
		return errors.New(fmt.Sprintf("error accessing Consul: %s", err))
	}

	return err
}

func (e *EventBus) Stop() {

}

func (e *EventBus) Consume(address string) *Consumer {
	return nil
}

func (e *EventBus) Publish(address string, message *Message) {

}

func (e *EventBus) Send(address string, message *Message) {

}

func (e *EventBus) handle(rawevent *api.UserEvent) {
	var msg Message
	err := json.Unmarshal(rawevent.Payload, &msg)
	if err != nil {
		e.Emit("error", fmt.Sprint("error parsing data for msg:", rawevent.ID))
	} else {
		msg.Id = rawevent.ID
		e.Emit("event", msg)
		e.Emit(msg.To, msg)
	}
}