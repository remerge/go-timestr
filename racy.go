// +build !race

package timestr

import (
	"sync"
	"time"
)

var ticker *time.Ticker
var stopped sync.WaitGroup
var done = make(chan bool)

func updateTicker() {
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

// Today returns the current time truncated to midnight
func Today() time.Time {
	return timeToday
}

// ISO8601 returns the ISO8601 representation of Now()
func ISO8601() string {
	return timeStrISO8601
}

// URLSafe return a URL-safe version of ISO8601()
func URLSafe() string {
	return timeStrURLSafe
}

// Stop stops the internal ticker and the cached values are not updated anymore
func Stop() {
	ticker.Stop()
	done <- true
	stopped.Wait()
}
