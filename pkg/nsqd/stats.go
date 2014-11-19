package nsqd

// stats with envelope.
type stats struct {
	Data *Stats
}

// Stats.
type Stats struct {
	Version string   `json:"version"`
	Health  string   `json:"health"`
	Topics  []*Topic `json:"topics"`
}

// Topic stats.
type Topic struct {
	Name          string     `json:"topic_name"`
	InFlightCount int64      `json:"in_flight_count"`
	DeferredCount int64      `json:"deferred_count"`
	MessageCount  int64      `json:"message_count"`
	RequeueCount  int64      `json:"requeue_count"`
	TimeoutCount  int64      `json:"timeout_count"`
	BackendDepth  int64      `json:"backend_depth"`
	Depth         int64      `json:"depth"`
	Paused        bool       `json:"paused"`
	Channels      []*Channel `json:"channels"`
}

// Channel stats.
type Channel struct {
	Name          string `json:"channel_name"`
	InFlightCount int64  `json:"in_flight_count"`
	DeferredCount int64  `json:"deferred_count"`
	MessageCount  int64  `json:"message_count"`
	RequeueCount  int64  `json:"requeue_count"`
	TimeoutCount  int64  `json:"timeout_count"`
	BackendDepth  int64  `json:"backend_depth"`
	Depth         int64  `json:"depth"`
	Paused        bool   `json:"paused"`
}
