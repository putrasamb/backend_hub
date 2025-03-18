package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"service-collection/internal/adapter/validator"
	"service-collection/internal/infrastructure/logger"
	httprequestmodel "service-collection/pkg/common/http/request/model"

	"github.com/labstack/echo/v4"
)

// LogController represents log controller
type LogController struct {
	Logger    *logger.Logger
	Validator *validator.CustomValidator
}

// HealthCheckController instantiate new log controller
func NewLogController(l *logger.Logger, v *validator.CustomValidator) *LogController {
	return &LogController{
		Logger:    l,
		Validator: v,
	}
}

func (ctrl *LogController) logError(err error, msg interface{}) {
	ctrl.Logger.Logger.WithError(err).Error(msg)
}

// Read reads log file
func (ctrl *LogController) Read(c echo.Context) error {
	request := &httprequestmodel.ReadLogRequest{}
	if err := c.Bind(request); err != nil {
		errMessage := "failed to parse read log request body"
		ctrl.logError(err, errMessage)
		return c.String(http.StatusBadRequest, errMessage)
	}

	if err := ctrl.Validator.Validate(request); err != nil {
		errMessage := "failed to validate request"
		validationErrors := ctrl.Validator.ParseValidationErrors(err)
		ctrl.logError(validationErrors, errMessage)
		return c.String(http.StatusBadRequest, errMessage)
	}

	if request.F == "" || request.F == ":f" {
		request.F = ctrl.Logger.FileName
	}

	fileAddress, err := ctrl.Logger.GetLogFilePath(request.F, request.Year, request.Month, request.Day)
	if err != nil {
		errMessage := "failed to find log file address"
		ctrl.Logger.Logger.Error()
		ctrl.logError(err, errMessage)
		return c.String(http.StatusNotFound, errMessage)
	}

	if request.N == 0 {
		src, err := os.ReadFile(*fileAddress)
		if err != nil {
			return c.String(http.StatusNotFound, "failed to find log file")
		}
		return c.String(http.StatusOK, string(src))
	}

	src, err := os.Open(*fileAddress)
	if err != nil {
		return c.String(http.StatusNotFound, "failed to open log file")
	}
	defer src.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := src.Stat()
	filesize := stat.Size()

	var lines []string

	for {
		cursor -= 1
		src.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		_, err := src.Read(char)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %v", err))
		}

		if char[0] == 10 || char[0] == 13 { // Newline or carriage return
			// Prepend the line to the slice of lines
			lines = append([]string{line}, lines...)
			line = ""

			if len(lines) >= request.N {
				break
			}
		} else {
			// Prepend the character to the line
			line = fmt.Sprintf("%s%s", string(char), line)
		}

		if cursor == -filesize { // Stop if at the beginning of the file
			// Append the last line if the beginning is reached
			if len(line) > 0 {
				lines = append([]string{line}, lines...)
			}
			break
		}
	}

	return c.String(http.StatusOK, strings.Join(lines, "\n"))
}
