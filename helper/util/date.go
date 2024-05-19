package util

import "time"

var (
	LayoutDefault  = "2006-01-02 15:04:05"
	LayoutDateOnly = "2006-01-02"
	LocalTimeZone  = "Asia/Ho_Chi_Minh"
	Loc, _         = time.LoadLocation(LocalTimeZone)
)

func TimeNow() time.Time {
	return time.Now().In(Loc)
}
