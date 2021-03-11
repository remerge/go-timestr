// +build !race

package timestr

import (
	"sync"
	"sync/atomic"
	"time"
)

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
			updateTimeStr()
		}
	}
}

func init() {
	updateTimeStr()
	ticker = time.NewTicker(1 * time.Second)
	go updateTicker()
}

// Now returns the current time with a 1-second precision
func Now() time.Time {
	return timeNow
}

// NowUTC returns the current time with a 1-second precision (in UTC TZ)
func NowUTC() time.Time {
	return timeNowUTC
}

// Today returns the current time truncated to midnight
func Today() time.Time {
	return timeToday
}

// Today returns the current time truncated to midnight (in UTC TZ)
func TodayUTC() time.Time {
	return timeTodayUTC
}

// ISO8601 returns the ISO8601 representation of Now()
func ISO8601() string {
	return timeStrISO8601
}

// ISO8601inUTC returns the ISO8601 representation of Now() (in UTC TZ)
func ISO8601inUTC() string {
	return timeStrISO8601inUTC
}

// URLSafe return a URL-safe version of ISO8601()
func URLSafe() string {
	return timeStrURLSafe
}

// URLSafeinUTC return a URL-safe version of ISO8601inUTC() (in UTC TZ)
func URLSafeinUTC() string {
	return timeStrURLSafeinUTC
}

// Stop stops the internal ticker and the cached values are not updated anymore
func Stop() {
	if !atomic.CompareAndSwapInt32(&started, 1, 0) {
		return
	}
	ticker.Stop()
	done <- true
	stopped.Wait()
}
