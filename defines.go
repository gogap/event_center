package event_center

type PushMode int32

const (
	SequencyMode           PushMode = 1
	SequencyAndWaitMode             = 2
	ConcurrencyMode                 = 3
	ConcurrencyAndWaitMode          = 4
)
