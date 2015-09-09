package event_center

type EventCenter interface {
	Name() string
	RegisterEvent(pushMode PushMode, eventNames ...string) (err error)
	Subscribe(eventName string, subscriber *Subscriber) (err error)
	Unsubscribe(eventName string, subscriberIds ...string) (err error)
	UnsubscribeAll(eventName string) (err error)
	PushEvent(eventName string, values ...interface{})
	ListEvents() []string
}

var (
	defaultEventCenter EventCenter = NewClassicEventCenter("default")
)

func RegisterEvent(pushMode PushMode, eventNames ...string) (err error) {
	err = defaultEventCenter.RegisterEvent(pushMode, eventNames...)
	return
}

func Subscribe(eventName string, subscriber *Subscriber) (err error) {
	err = defaultEventCenter.Subscribe(eventName, subscriber)
	return
}

func Unsubscribe(eventName string, subscriberIds ...string) (err error) {
	err = defaultEventCenter.Unsubscribe(eventName, subscriberIds...)
	return
}

func PushEvent(eventName string, values ...interface{}) {
	defaultEventCenter.PushEvent(eventName, values...)
	return
}

func UnsubscribeAll(eventName string) (err error) {
	err = defaultEventCenter.UnsubscribeAll(eventName)
	return
}

func ListEvents() []string {
	return defaultEventCenter.ListEvents()
}
