package logger

import (
	"encoding/json"
	"time"
)

// ActivityLog represents a single activity log entry.
type ActivityLog struct {
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Activity     string    `json:"activity"`
	Module       string    `json:"module"`
	ActivityTime time.Time `json:"activity_time"`
	IPAddress    string    `json:"ip_address,omitempty"`
	DeviceInfo   string    `json:"device_info,omitempty"`
	Location     string    `json:"location,omitempty"`
	Remarks      string    `json:"remarks,omitempty"`
}

// Json returns activity log's json representation
func (m *ActivityLog) Json() ([]byte, error) {
	return json.Marshal(m)
}

// String returns activity log's json string reprensentation
func (m *ActivityLog) String() (*string, error) {
	b, err := m.Json()
	if err != nil {
		return nil, err
	}
	str := string(b)
	return &str, nil
}

// ActivityLogger represents activity logger
type ActivityLogger struct {
	LogPublisher[*ActivityLog]
}
