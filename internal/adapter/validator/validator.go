package validator

import (
	"fmt"

	model "service-collection/internal/adapter/validator/model"

	"github.com/go-playground/validator/v10"
)

type ValidatorInterface interface {
	Validate(i interface{}) error
	ParseValidationErrors(err error) error
}

// CustomValidator represents echo's custom validator
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validate struct.
// This method adds chain function functionality to *echo.Context instance.
// So we can use it like: *echo.Context.Validate(i interface{})
func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// ParseValidationErrors parse validation errors into new ValidationError
func (v *CustomValidator) ParseValidationErrors(err error) error {
	var errs model.ValidationErrors

	if err == nil {
		return errs
	}
	for _, err := range err.(validator.ValidationErrors) {
		e := model.ValidationError{
			Namespace:       err.Namespace(),
			Field:           err.Field(),
			StructNamespace: err.StructNamespace(),
			StructField:     err.StructField(),
			Tag:             err.Tag(),
			ActualTag:       err.ActualTag(),
			Kind:            fmt.Sprintf("%v", err.Kind()),
			Type:            fmt.Sprintf("%v", err.Type()),
			Value:           fmt.Sprintf("%v", err.Value()),
			Param:           err.Param(),
			Message:         err.Error(),
		}

		errs.Errors = append(errs.Errors, e)
	}

	return errs
}

func NewValidator() *CustomValidator {
	v := validator.New()
	registerCommonRules(v)
	return &CustomValidator{
		Validator: v,
	}
}

func registerCommonRules(v *validator.Validate) {
	v.RegisterValidation("lt_today", dateOlderThanToday)
	v.RegisterValidation("lte_today", dateOlderOrEqualThanToday)
	v.RegisterValidation("gt_today", dateGreaterThanToday)
	v.RegisterValidation("gte_today", dateGreaterOrEqualThanToday)
}
