package logger

import (
	"encoding/json"
	"time"
)

// AuditLog represents a single audit log entry.
type AuditLog struct {
	Module     string    `json:"module"`
	ActionType string    `json:"actionType"`
	SearchKey  string    `json:"searchKey"`
	Before     string    `json:"before"`
	After      string    `json:"after"`
	ActionBy   string    `json:"actionBy"`
	ActionTime time.Time `json:"timestamp"`
}

// Json returns audit log's json representation
func (m *AuditLog) Json() ([]byte, error) {
	return json.Marshal(m)
}

// String returns audit log's json string reprensentation
func (m *AuditLog) String() (*string, error) {
	b, err := m.Json()
	if err != nil {
		return nil, err
	}
	str := string(b)
	return &str, nil
}

// AuditLogger represents audit logger
type AuditLogger struct {
	LogPublisher[*AuditLog]
}
