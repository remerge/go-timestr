// +build race

package timestr

import (
	"sync"
	"time"
)

var timeStrMutex = sync.RWMutex{}

func init() {
	UpdateTimeStr()
	go func() {
		for range time.NewTicker(1 * time.Second).C {
			timeStrMutex.Lock()
			UpdateTimeStr()
			timeStrMutex.Unlock()
		}
	}()
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
