package naming

import (
	"fmt"
	"strings"
	"time"
)

type NamingType int

const (
	Hash NamingType = iota
	UUID
	Series
)

// String - Creating common behavior - give the type a String function
func (n NamingType) String() string {
	return [...]string{"Hash", "UUID", "Series"}[n]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (n NamingType) EnumIndex() int {
	return int(n)
}

type NamingHash struct {
	Length int
}

type NamingUUID struct {
}

// Naming represents generic function to generate document naming
type NamingSeries struct {
	Format string    // format separated by (.). Use * for using (.)
	Number int       // an integer to parse to "#" format
	Time   time.Time // given a timestamp. Default to time.now() on nil
}

// Parse generate naming series based on the given format and values
func (opt *NamingSeries) Parse() *string {

	n := ""

	if opt.Time.IsZero() {
		opt.Time = time.Now()
	}

	formats := strings.Split(opt.Format, ".")
	series_set := false

	for _, format := range formats {
		if format == "" {
			continue
		}

		part := ""

		if strings.HasPrefix(format, "#") {
			if !series_set {
				digits := len(format)
				part = fmt.Sprintf("%0*d", digits, opt.Number)
				series_set = true
			}
		} else if format == "YY" {
			part = opt.Time.Format("06")
		} else if format == "YYYY" {
			part = opt.Time.Format("2006")
		} else if format == "MM" {
			part = opt.Time.Format("01")
		} else if format == "DD" {
			part = opt.Time.Format("02")
		} else if format == "WW" {
			_, w := opt.Time.ISOWeek()

			part = fmt.Sprintf("%02d", w)
		} else {
			part = format
		}

		n += part
	}

	return &n
}
