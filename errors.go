package event_center

import (
	"github.com/gogap/errors"
)

const EVENT_CENTER_ERR_NS = "EVENT_CENTER"

var (
	ERR_SUBSCRIBER_NOT_EXIST               = errors.TN(EVENT_CENTER_ERR_NS, 1, "subscriber {{.id}} not exist")
	ERR_COULD_NOT_GET_ANONYMOUS_SUBSCRIBER = errors.TN(EVENT_CENTER_ERR_NS, 2, "could not get subscriber anonymous by name")
	ERR_INVALIDATE_INDEX_RANGE             = errors.TN(EVENT_CENTER_ERR_NS, 3, "invalidate index range")
	ERR_SUBSCRIBER_ID_IS_EMPTY             = errors.TN(EVENT_CENTER_ERR_NS, 4, "subscriber id is empty")
	ERR_SUBSCRIBER_HANDLER_IS_NIL          = errors.TN(EVENT_CENTER_ERR_NS, 5, "subscriber handler is nil, index: {{.index}}, name: {{.id}}")
	ERR_SUBSCRIBER_ID_GENERATE_FAILED      = errors.TN(EVENT_CENTER_ERR_NS, 6, "generate subscriber id failed, err: {{.err}}")
	ERR_SUBSCRIBER_IS_NIL                  = errors.TN(EVENT_CENTER_ERR_NS, 7, "subscriber is nil")

	ERR_EVENT_ALREADY_REGISTERED = errors.TN(EVENT_CENTER_ERR_NS, 8, "event already registered, name: {{.name}}")
	ERR_EVENT_NAME_IS_EMPTY      = errors.TN(EVENT_CENTER_ERR_NS, 9, "event name is empty")
	ERR_EVENT_NOT_EXIST          = errors.TN(EVENT_CENTER_ERR_NS, 10, "event not exist, name: {{.name}}")

	ERR_EVENT_CENTER_NAME_IS_EMPTY = errors.TN(EVENT_CENTER_ERR_NS, 11, "event center name is empty")
)
