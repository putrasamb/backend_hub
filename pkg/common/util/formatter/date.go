package util

import "time"

// YYYY-MM-DD
func DateFormatter(date string) string {
	newFormatDate, _ := time.Parse("2006-01-02T15:04:05-07:00", date)
	return newFormatDate.Format("2006-01-02")
}

type CustomDate struct {
	time.Time
}

// MarshalJSON formats the date as "2006-01-02".
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cd.Time.Format("2006-01-02") + `"`), nil
}

// UnmarshalJSON parses the date from "2006-01-02".
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	str := string(data)
	parsedTime, err := time.Parse(`"2006-01-02"`, str)
	if err != nil {
		return err
	}
	cd.Time = parsedTime
	return nil
}

func (cd CustomDate) String() string {
	return cd.Time.Format("2006-01-02")
}
