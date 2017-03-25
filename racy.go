// +build !race

package timestr

import "time"

func init() {
	UpdateTimeStr()
	go func() {
		for range time.NewTicker(1 * time.Second).C {
			UpdateTimeStr()
		}
	}()
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
