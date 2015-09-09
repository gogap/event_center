Event Center
============

A Event-Driven library for decoupling each component


#### EventCenter Interface

```go
type EventCenter interface {
	Name() string
	RegisterEvent(pushMode PushMode, eventNames ...string) (err error)
	Subscribe(eventName string, subscriber *Subscriber) (err error)
	Unsubscribe(eventName string, subscriberIds ...string) (err error)
	UnsubscribeAll(eventName string) (err error)
	PushEvent(eventName string, values ...interface{})
	ListEvents() []string
}
```


Usage:

#### Create An Event Center

```go
import (
	"github.com/gogap/event_center"
)

...

EventCenter event_center.EventCenter = event_center.NewClassicEventCenter("GoGapEventCenter")

```

Or you can just use the default EventCenter like as following

```go

event_center.RegisterEvent(...)
...
...

```


#### Register Event

```go
const (
	EVENT_CMD_STOP			= "EVENT_CMD_STOP"
	EVENT_RECEIVER_STOPPED	= "EVENT_RECEIVER_STOPPED"
)

...
EventCenter.RegisterEvent(event_center.ConcurrencyAndWaitMode,
		EVENT_CMD_STOP,
		EVENT_RECEIVER_STOPPED,
	)
...
```

#### Subscribe Event

```go

...
stopSubscriber := event_center.NewSubscriber(func(eventName string, values ...interface{}) {
		if !isStoped {
			isStoped = true

			EventCenter.PushEvent(EVENT_RECEIVER_STOPPED, ....)
		}
	})

EventCenter.Subscribe(EVENT_CMD_STOP, stopSubscriber)
...

```


#### Push Event

```go
...
EventCenter.PushEvent(EVENT_RECEIVER_STOPPED, ....)
...
```