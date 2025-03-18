package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Logger represents Logger
type Logger struct {
	Logger         *logrus.Logger
	ActivityLogger *ActivityLogger
	AuditLogger    *AuditLogger
	Output         string
	FileName       string
}

// LogFields represents log data structure
type LogFields struct {
	Time      time.Time `json:"time,omitempty"`
	Level     string    `json:"level,omitempty"`
	Message   string    `json:"msg,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
}

func (l *Logger) validate() error {

	if l.Output == "" {
		return errors.New("output directory is mandatory")
	}
	if l.FileName == "" {
		l.FileName = "logs"
	}
	return nil
}

func (l *Logger) Format(entry *logrus.Entry) ([]byte, error) {
	var buf bytes.Buffer

	// Create a map to hold the log entry fields
	logData := make(map[string]interface{})
	logData["time"] = entry.Time.Format(time.RFC3339)
	logData["level"] = entry.Level.String()
	logData["msg"] = entry.Message

	// Include additional fields
	for key, value := range entry.Data {
		logData[key] = value
	}

	// Marshal the map into pretty JSON
	prettyJSON, err := json.MarshalIndent(logData, "", " ")
	if err != nil {
		return nil, err
	}
	buf.Write(prettyJSON)
	buf.WriteByte('\n') // Add a newline for better readability

	return buf.Bytes(), nil
}

// Setup setup logger format and log locations.
func (l *Logger) Setup() error {

	if err := l.validate(); err != nil {
		return err
	}

	now := time.Now()

	fullpath, err := l.GetLogFilePath(l.FileName, now.Year(), now.Month(), now.Day())

	if err != nil {
		return err
	}
	file, err := os.OpenFile(*fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "failed to open log destination")
	}
	l.Logger.SetOutput(file)
	l.Logger.SetFormatter(l)
	return nil
}

// GetLogFilePath generates the path for the log file
func (l *Logger) GetLogFilePath(filename string, y int, m time.Month, d int) (*string, error) {
	dir := fmt.Sprintf("%s/%d/%d/%d", l.Output, y, m, d)
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find absolute log directory")
	}

	os.MkdirAll(abs, os.ModePerm) // Create directories if they don't exist
	path := filepath.Join(abs, fmt.Sprintf("%s.log", filename))
	return &path, nil
}
