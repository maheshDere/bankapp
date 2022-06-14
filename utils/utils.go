package utils

import "time"

func ParseStringToTime(s string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", s)
	return
}
