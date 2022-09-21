package vo

import "time"

type Date string

const dateFormat = "2006-01-02"

func (d Date) GetTime() (time.Time, error) {
	if string(d) == "" {
		return time.Time{}, nil
	}
	return time.Parse(dateFormat, string(d))
}
