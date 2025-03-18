package util

import "time"

func GetDateOrDefault(date *time.Time, defaultDays int) time.Time {
	if date == nil {
		return time.Now().AddDate(0, 0, defaultDays)
	}
	return *date
}

func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format("2006-01-02")
}
