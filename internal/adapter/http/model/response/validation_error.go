package model

// ResponseValidationError represents a single error response
type ResponseValidationError struct {
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

// ResponseValidationErrors represents slice of errors respose
type ResponseValidationErrors struct {
	Errors []ResponseValidationError `json:"errors,omitempty"`
}
