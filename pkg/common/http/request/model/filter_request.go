package httprequestmodel

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
)

// QueryString represents querystring default format
type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type ListRequest struct {
	Page                 int    `query:"page"`
	PerPage              int    `query:"per_page"`
	FiltersStringEncoded string `query:"filters"`
	Filters              *[]Filter
}

// parseFunc decode query filter base64 URL encoded
func parseFunc(b64 string) (*[]Filter, error) {
	var filters []Filter

	if b64 == "" {
		return &filters, nil
	}

	decoded, err := base64.URLEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(decoded, &filters); err != nil {
		return nil, errors.Wrap(err, "invalid filter format")
	}
	return &filters, nil
}

// DecodeFilters returns decoded query filters
func (r *ListRequest) DecodeFilters() error {

	filters, err := parseFunc(r.FiltersStringEncoded)
	if err != nil {
		return err
	}
	r.Filters = filters
	return nil
}
