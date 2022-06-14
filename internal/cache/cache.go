package cache

import (
	"time"
)

type Cache interface {
	Ping() error
	Get(string) (string, error)
	Set(string, interface{}, time.Duration) (string, error)
}
