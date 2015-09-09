package event_center

import (
	"sync"

	"github.com/gogap/errors"
)

type ClassicEventCenter struct {
	eventGroupLocker sync.Mutex

	name        string
	eventGroups map[string]*SubscriberGroup
}

func NewClassicEventCenter(name string) (eventCenter EventCenter) {
	if name == "" {
		err := ERR_EVENT_CENTER_NAME_IS_EMPTY.New()
		panic(err)
	}

	eventCenter = &ClassicEventCenter{
		name:        name,
		eventGroups: make(map[string]*SubscriberGroup),
	}

	return
}

func (p *ClassicEventCenter) Name() string {
	return p.name
}

func (p *ClassicEventCenter) RegisterEvent(pushMode PushMode, eventNames ...string) (err error) {
	p.eventGroupLocker.Lock()
	defer p.eventGroupLocker.Unlock()

	if eventNames == nil || len(eventNames) == 0 {
		return
	}

	for _, eventName := range eventNames {
		if eventName == "" {
			err = ERR_EVENT_ALREADY_REGISTERED.New()
			return
		}

		if _, exist := p.eventGroups[eventName]; exist {
			err = ERR_EVENT_NAME_IS_EMPTY.New(errors.Params{"name": eventName})
			return
		}

		p.eventGroups[eventName] = NewSubscriberGroup(pushMode)
	}

	return
}

func (p *ClassicEventCenter) Subscribe(eventName string, subscriber *Subscriber) (err error) {
	p.eventGroupLocker.Lock()
	defer p.eventGroupLocker.Unlock()

	if eventName == "" {
		err = ERR_EVENT_NAME_IS_EMPTY.New()
		return
	}

	if eventGroup, exist := p.eventGroups[eventName]; !exist {
		err = ERR_EVENT_NOT_EXIST.New(errors.Params{"name": eventName})
		return
	} else {
		err = eventGroup.AddSubscriber(subscriber)
	}

	return
}

func (p *ClassicEventCenter) Unsubscribe(eventName string, subscriberNames ...string) (err error) {
	p.eventGroupLocker.Lock()
	defer p.eventGroupLocker.Unlock()

	if eventName == "" {
		err = ERR_EVENT_NAME_IS_EMPTY.New()
		return
	}

	if eventGroup, exist := p.eventGroups[eventName]; !exist {
		err = ERR_EVENT_NOT_EXIST.New(errors.Params{"name": eventName})
		return
	} else {
		err = eventGroup.RemoveSubscriber(subscriberNames...)
	}

	return
}

func (p *ClassicEventCenter) UnsubscribeAll(eventName string) (err error) {
	p.eventGroupLocker.Lock()
	defer p.eventGroupLocker.Unlock()

	if eventName == "" {
		err = ERR_EVENT_NAME_IS_EMPTY.New()
		return
	}

	if eventGroup, exist := p.eventGroups[eventName]; !exist {
		err = ERR_EVENT_NOT_EXIST.New(errors.Params{"name": eventName})
		return
	} else {
		eventGroup.ClearSubscriber()
	}

	return
}

func (p *ClassicEventCenter) PushEvent(eventName string, values ...interface{}) {
	if eventName == "" {
		return
	}

	if eventGroup, exist := p.eventGroups[eventName]; !exist {
		return
	} else {
		eventGroup.PushEvent(eventName, values...)
	}

	return
}

func (p *ClassicEventCenter) ListEvents() (events []string) {
	for event, _ := range p.eventGroups {
		events = append(events, event)
	}
	return
}
