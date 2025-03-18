package domain

import "fmt"

type ValidationError struct {
	Code    int
	Message string
	Err     error
}

func NewValidationError(code int, message string, err error) *ValidationError {
	return &ValidationError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *ValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Error: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type FieldValidationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Field   string `json:"field"`
}

type FieldValidationErrors struct {
	Errors []FieldValidationError `json:"errors"`
}

func NewFieldValidationErrors() *FieldValidationErrors {
	return &FieldValidationErrors{}
}

func (f *FieldValidationErrors) AddErrorDetail(code int, message, errorDetail, field string) {
	f.Errors = append(f.Errors, FieldValidationError{
		Code:    code,
		Message: message,
		Error:   errorDetail,
		Field:   field,
	})
}

func (f *FieldValidationErrors) Error() string {
	return fmt.Sprintf("Validation Errors: %v", f.GetErrorMessages())
}

func (f *FieldValidationErrors) GetErrorMessages() []string {
	var errorMessages []string
	for _, err := range f.Errors {
		errorMessages = append(errorMessages, fmt.Sprintf("%s (Field: %s)", err.Message, err.Field))
	}
	return errorMessages
}

func (f *FieldValidationErrors) HasErrors() bool {
	return len(f.Errors) > 0
}
