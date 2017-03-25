package timestr

import "time"

var timeNow time.Time
var timeToday time.Time
var timeStrISO8601 = ""
var timeStrURLSafe = ""

func updateTimeStr() {
	timeNow = time.Now()
	d := time.Duration(-timeNow.Hour()) * time.Hour
	timeToday = timeNow.Truncate(time.Hour).Add(d)
	timeStrISO8601 = timeNow.Format("2006-01-02T15:04:05Z")
	timeStrURLSafe = timeNow.Format("2006-01-02T15-04-05Z")
}
