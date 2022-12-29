package singleton

import "sync"

type single_3 map[string]string

var (
	once       sync.Once
	instance_3 single_3
)

func New() single_3 {
	once.Do(func() {
		instance_3 = make(single_3)
	})
	return instance_3
}
