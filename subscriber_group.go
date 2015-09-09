package event_center

import (
	"sort"
	"sync"

	"github.com/gogap/errors"
)

type Subscribers []*Subscriber

func (p Subscribers) Len() int {
	return len(p)
}

func (p Subscribers) Less(i, j int) bool {
	return p[i].Weight > p[j].Weight
}

func (p Subscribers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type SubscriberGroup struct {
	subscriberLocker sync.Mutex
	pushMode         PushMode
	subscribers      Subscribers
	subscribersIndex map[string]*Subscriber
}

func NewSubscriberGroup(pushMode PushMode) *SubscriberGroup {
	return &SubscriberGroup{
		pushMode:         pushMode,
		subscribersIndex: make(map[string]*Subscriber),
	}
}

func (p *SubscriberGroup) AddSubscriber(subscribers ...*Subscriber) (err error) {
	p.subscriberLocker.Lock()
	defer p.subscriberLocker.Unlock()

	if subscribers == nil || len(subscribers) == 0 {
		return
	}

	for index, subscriber := range subscribers {
		if subscriber == nil {
			err = ERR_SUBSCRIBER_IS_NIL.New(errors.Params{"index": index, "id": subscriber.Id})
			return
		}

		if subscriber.Handler == nil {
			err = ERR_SUBSCRIBER_HANDLER_IS_NIL.New(errors.Params{"index": index, "id": subscriber.Id})
			return
		}

		if subscriber.id == "" {
			err = ERR_SUBSCRIBER_ID_IS_EMPTY.New()
		}
	}

	p.clearDiedSubscriber()

	for _, subscriber := range subscribers {
		p.subscribersIndex[subscriber.id] = subscriber
	}

	p.subscribers = append(p.subscribers, subscribers...)

	sort.Sort(p.subscribers)

	return
}

func (p *SubscriberGroup) clearDiedSubscriber() {
	for key, subscriber := range p.subscribersIndex {
		if subscriber == nil {
			delete(p.subscribersIndex, key)
		}
	}

	newSubscribers := make([]*Subscriber, 0)

	for _, subscriber := range p.subscribers {
		if subscriber != nil {
			newSubscribers = append(newSubscribers, subscriber)
		}
	}

	p.subscribers = newSubscribers
}

func (p *SubscriberGroup) SubscriberNameList() (subscriberNameList []string, err error) {
	for _, subscriber := range p.subscribers {
		subscriberNameList = append(subscriberNameList, subscriber.id)
	}

	return
}

func (p *SubscriberGroup) GetSubscriber(subscriberIds ...string) (subscriber []*Subscriber, err error) {
	p.subscriberLocker.Lock()
	defer p.subscriberLocker.Unlock()

	if subscriberIds == nil || len(subscriberIds) == 0 {
		return
	}

	subscriber = make([]*Subscriber, len(subscriberIds))
	for _, id := range subscriberIds {
		for _, suber := range p.subscribers {
			if suber != nil && suber.id == id {
				subscriber = append(subscriber, suber)
			}
		}
	}

	return
}

func (p *SubscriberGroup) RemoveSubscriber(subscriberIds ...string) (err error) {
	p.subscriberLocker.Lock()
	defer p.subscriberLocker.Unlock()

	if subscriberIds == nil || len(subscriberIds) == 0 {
		return
	}

	for _, id := range subscriberIds {
		if id == "" {
			err = ERR_SUBSCRIBER_ID_IS_EMPTY.New()
			return
		}

		if _, exist := p.subscribersIndex[id]; !exist {
			err = ERR_SUBSCRIBER_NOT_EXIST.New(errors.Params{"id": id})
			return
		}
	}

	newSubscribers := make([]*Subscriber, len(p.subscribers)-len(subscriberIds))
	for _, id := range subscriberIds {
		for _, subscriber := range p.subscribers {
			if subscriber != nil && subscriber.id != id {
				newSubscribers = append(newSubscribers, subscriber)
			} else {
				delete(p.subscribersIndex, id)
			}
		}
	}

	p.subscribers = newSubscribers

	return
}

func (p *SubscriberGroup) ClearSubscriber() {
	p.subscriberLocker.Lock()
	defer p.subscriberLocker.Unlock()

	p.subscribers = make([]*Subscriber, 0)
	p.subscribersIndex = make(map[string]*Subscriber)
}

func (p *SubscriberGroup) PushEvent(eventName string, values ...interface{}) {
	if p.pushMode == ConcurrencyMode {
		for _, subscriber := range p.subscribers {
			go subscriber.Handler(eventName, values...)
		}
	} else if p.pushMode == ConcurrencyAndWaitMode {
		wg := sync.WaitGroup{}

		funcHandler := func(handler SubscriberHandler, eventName string, v ...interface{}) {
			defer wg.Done()
			handler(eventName, v...)
		}

		for _, subscriber := range p.subscribers {
			wg.Add(1)
			go funcHandler(subscriber.Handler, eventName, values...)
		}
		wg.Wait()
	} else if p.pushMode == SequencyMode {
		go func(eventName string, values ...interface{}) {
			for _, subscriber := range p.subscribers {
				subscriber.Handler(eventName, values...)
			}
		}(eventName, values...)
	} else if p.pushMode == SequencyAndWaitMode {
		for _, subscriber := range p.subscribers {
			subscriber.Handler(eventName, values...)
		}
	}
}
