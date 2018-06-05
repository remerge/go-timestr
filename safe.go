// +build race

package timestr

import (
	"sync"
	"sync/atomic"
	"time"
)

var timeStrMutex = sync.RWMutex{}
var ticker *time.Ticker
var stopped sync.WaitGroup
var done = make(chan bool)
var started int32

func updateTicker() {
	if !atomic.CompareAndSwapInt32(&started, 0, 1) {
		return
	}
	stopped.Add(1)
	defer stopped.Done()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			timeStrMutex.Lock()
			updateTimeStr()
			timeStrMutex.Unlock()
		}
	}
}

func init() {
	updateTimeStr()
	ticker = time.NewTicker(1 * time.Second)
	go updateTicker()
}

func Now() time.Time {
	timeStrMutex.RLock()
	defer timeStrMutex.RUnlock()
	return timeNow
}

func Today() time.Time {
	timeStrMutex.RLock()
	defer timeStrMutex.RUnlock()
	return timeToday
}

func ISO8601() string {
	timeStrMutex.RLock()
	defer timeStrMutex.RUnlock()
	return timeStrISO8601
}

func URLSafe() string {
	timeStrMutex.RLock()
	defer timeStrMutex.RUnlock()
	return timeStrURLSafe
}

func Stop() {
	if !atomic.CompareAndSwapInt32(&started, 1, 2) {
		return
	}
	ticker.Stop()
	done <- true
	stopped.Wait()
}
