package scheduler

import (
	"time"
)

func Schedule(fn func(), delay time.Duration) chan bool {
	stopScheduler := make(chan bool)

	go func() {
		for {
			fn()
			select {
			case <-time.After(delay):
			case <-stopScheduler:
				return
			}
		}
	}()

	return stopScheduler
}
