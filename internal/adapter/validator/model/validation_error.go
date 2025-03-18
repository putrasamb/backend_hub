package model

import (
	"encoding/json"

	model "backend_hub/internal/adapter/http/model/response"
)

type ValidationError struct {
	Namespace       string `json:"namespace,omitempty"` // can differ when a custom TagNameFunc is registered or
	Field           string `json:"field,omitempty"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace,omitempty"`
	StructField     string `json:"structField,omitempty"`
	Tag             string `json:"tag,omitempty"`
	ActualTag       string `json:"actualTag,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Type            string `json:"type,omitempty"`
	Value           string `json:"value,omitempty"`
	Param           string `json:"param,omitempty"`
	Message         string `json:"message,omitempty"`
}

func (e ValidationError) Error() string {
	result, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		panic(err)
	}
	return string(result)
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors,omitempty"`
}

func (errs ValidationErrors) Error() string {
	result, err := json.MarshalIndent(errs, "", " ")
	if err != nil {
		panic(err)
	}
	return string(result)
}

// ToErrorResponse converts error type into response struct
func (errs ValidationErrors) ToResponseErrors() []model.ResponseValidationError {
	var response []model.ResponseValidationError
	for _, err := range errs.Errors {
		e := &model.ResponseValidationError{
			Namespace:       err.Namespace,
			Field:           err.Field,
			StructNamespace: err.StructNamespace,
			StructField:     err.StructField,
			Tag:             err.Tag,
			ActualTag:       err.ActualTag,
			Kind:            err.Kind,
			Type:            err.Type,
			Value:           err.Value,
			Param:           err.Param,
			Message:         err.Message,
		}

		response = append(response, *e)
	}
	return response
}
