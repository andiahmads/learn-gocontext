package pengenalancontext

import "time"

type Context interface {
	Deadline() (deadline time.Time)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
