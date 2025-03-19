package httpresponse

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoResponse struct{}

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Response
	Error interface{} `json:"error,omitempty"`
}

type DataResponse[T any] struct {
	Response
	Data T `json:"data,omitempty"`
}

type PaginatedReponse[T any] struct {
	Data        []T   `json:"data"`
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page,omitempty"`
	PerPage     int   `json:"per_page,omitempty"`
	Page        int   `json:"page,omitempty"`
}

func (r *PaginatedReponse[T]) EchoJsonResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, r)
}

func (r *ErrorResponse) EchoJsonResponse(c echo.Context) error {
	return c.JSON(r.Code, r)
}

func (r *Response) EchoJsonResponse(c echo.Context) error {
	return c.JSON(r.Code, r)
}

func (r *Response) Error() string {
	return r.Message
}

func (r *Response) GetStatusCode() int {
	return r.Code
}

func NewHTTPError(code int, message string) error {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func NewErrorResponse(code int, message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Code:    code,
			Message: message,
		},
		Error: err.Error(),
	}
}
