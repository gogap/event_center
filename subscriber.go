package event_center

import (
	"github.com/nu7hatch/gouuid"
)

type SubscriberHandler func(eventName string, values ...interface{})

type Subscriber struct {
	id      string
	Handler SubscriberHandler
	Weight  int64
}

func (p *Subscriber) Id() string {
	return p.id
}

func NewSubscriber(handler SubscriberHandler) *Subscriber {
	subscriberId := ""
	if id, e := uuid.NewV4(); e != nil {
		return nil
	} else {
		subscriberId = id.String()
	}

	return &Subscriber{
		id:      subscriberId,
		Handler: handler,
		Weight:  0,
	}
}

func (p *Subscriber) SetWeight(weight int64) *Subscriber {
	p.Weight = weight
	return p
}
