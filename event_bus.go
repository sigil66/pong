package pong

import (
	"github.com/chuckpreslar/emission"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"fmt"
	"encoding/json"
)

const DefaultConsulUri = "http://localhost:8500"

type EventBus struct {
	*emission.Emitter

	Shutdown chan int

	consulUri string
	consulClient *api.Client
	consulWatch  *watch.Plan

	seen uint64
}

func NewEventBus(consulUri string) *EventBus {
	eb := &EventBus{Emitter: emission.NewEmitter(), Shutdown: make(chan int, 1)}
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
	e.On("rawmessage", e.handle)

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
				e.Emit("rawmessage", event)
				e.seen = event.LTime
			}
		}
	}

	go func() {
		if err := e.consulWatch.Run(e.consulUri); err != nil {
			e.Emit("error", fmt.Sprintf("error accessing Consul: %s", err))
		}

		e.Shutdown <- 0

	}()

	return err
}

func (e *EventBus) Stop() {
	e.consulWatch.Stop()
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
		e.Emit("error", fmt.Sprint("error parsing data for message:", rawevent.ID))
	} else {
		msg.Id = rawevent.ID
		e.Emit("message", msg)
		e.Emit(msg.To, msg)
	}
}