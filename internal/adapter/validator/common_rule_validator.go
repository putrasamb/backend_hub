package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// dateOlderThanToday validate if field is older than today.
func dateOlderThanToday(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)

	if !ok {
		return false
	}

	// Get the current date (start of today)
	today := time.Now().Truncate(24 * time.Hour)

	// Check if the given time is before today
	return t.Before(today)
}

// dateOlderOrEqualThanToday validate if field is older or equal than today.
func dateOlderOrEqualThanToday(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)

	if !ok {
		return false
	}

	// Get the current date (start of today)
	today := time.Now().Truncate(24 * time.Hour)
	return !t.After(today)
}

// dateGreaterThanToday validate if field is older or equal than today.
func dateGreaterThanToday(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)

	if !ok {
		return false
	}

	// Get the current date (start of today)
	today := time.Now().Truncate(24 * time.Hour)
	return t.After(today)
}

// dateGreaterOrEqualThanToday validate if field is older or equal than today.
func dateGreaterOrEqualThanToday(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)

	if !ok {
		return false
	}
	// Get the current date (start of today)
	today := time.Now().Truncate(24 * time.Hour)
	return !t.Before(today)
}
