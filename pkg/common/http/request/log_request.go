package httprequest

import "time"

type ReadLogRequest struct {
	Year  int        `param:"year" query:"year" validate:"required"`
	Month time.Month `param:"month" query:"month" validate:"required"`
	Day   int        `param:"day" query:"day" validate:"required"`
	N     int        `query:"n"`           // number of line
	F     string     `param:"f" query:"f"` // filename
}
