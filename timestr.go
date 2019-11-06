package timestr

import "time"

var timeNow time.Time
var timeNowUTC time.Time
var timeToday time.Time
var timeStrISO8601 = ""
var timeStrURLSafe = ""
var timeStrISO8601inUTC = ""
var timeStrURLSafeinUTC = ""

func updateTimeStr() {
	SetTimeStr(time.Now())
}

func SetTimeStr(t time.Time) {
	timeNow = t
	timeNowUTC = t.UTC()
	d := time.Duration(-t.Hour()) * time.Hour
	timeToday = t.Truncate(time.Hour).Add(d)
	timeStrISO8601 = t.Format("2006-01-02T15:04:05Z07:00")
	timeStrISO8601inUTC = timeNowUTC.Format("2006-01-02T15:04:05Z07:00")
	timeStrURLSafe = t.Format("2006-01-02T15-04-05Z0700")
	timeStrURLSafeinUTC = timeNowUTC.Format("2006-01-02T15-04-05Z0700")
}
