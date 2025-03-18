package httprequest

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type Sort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"` // asc or desc
}

type FilteredRequestInterface interface {
	DecodeFilters() error
	GetFilters() *[]Filter
	DecodeSort() error
	GetSort() *[]Sort
}

// QueryString represents querystring default format
type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type FilteredRequest struct {
	FiltersStringEncoded string `query:"filters"`
	Filters              *[]Filter
	SortStringEncoded    string `query:"sort"`
	Sort                 *[]Sort
}

type ListRequest struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
	FilteredRequest
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

// parseSortFunc decode query sort base64 URL encoded
func parseSortFunc(b64 string) (*[]Sort, error) {
	var sorts []Sort

	if b64 == "" {
		return &sorts, nil
	}

	decoded, err := base64.URLEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(decoded, &sorts); err != nil {
		return nil, errors.Wrap(err, "invalid sort format")
	}

	// Validate sort direction
	for i, sort := range sorts {
		sorts[i].Direction = strings.ToLower(sort.Direction)
		if sort.Direction != "asc" && sort.Direction != "desc" {
			sorts[i].Direction = "asc" // default to ascending if invalid
		}
	}

	return &sorts, nil
}

func (r *FilteredRequest) GetFilters() *[]Filter {
	return r.Filters
}

func (r *FilteredRequest) GetSort() *[]Sort {
	return r.Sort
}

// DecodeFilters returns decoded query filters
func (r *FilteredRequest) DecodeFilters() error {
	filters, err := parseFunc(r.FiltersStringEncoded)
	if err != nil {
		return err
	}
	r.Filters = filters
	return nil
}

// DecodeSort returns decoded query sorts
func (r *FilteredRequest) DecodeSort() error {
	sorts, err := parseSortFunc(r.SortStringEncoded)
	if err != nil {
		return err
	}
	r.Sort = sorts
	return nil
}
