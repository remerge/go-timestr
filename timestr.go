package timestr

import "time"

var timeNow time.Time
var timeToday time.Time
var timeStrISO8601 = ""
var timeStrURLSafe = ""

func updateTimeStr() {
	SetTimeStr(time.Now())
}

func SetTimeStr(t time.Time) {
	timeNow = t
	d := time.Duration(-t.Hour()) * time.Hour
	timeToday = t.Truncate(time.Hour).Add(d)
	timeStrISO8601 = t.Format("2006-01-02T15:04:05Z")
	timeStrURLSafe = t.Format("2006-01-02T15-04-05Z")
}
