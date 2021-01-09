package logger

import "time"

type (
	// Key context
	Key int
)

// LogKey ...
const LogKey = Key(48)

// Data logging
type Data struct {
	RequestID  string                 `json:"RequestID"`
	TimeStart  time.Time              `json:"TimeStart"`
	ID         string                 `json:"ID"`
	StatusCode int                    `json:"StatusCode"`
	Service    string                 `json:"Service"`
	Host       string                 `json:"Host"`
	Endpoint   string                 `json:"Endpoint"`
	Method     string                 `json:"Method"`
	Body       map[string]interface{} `json:"Body"`
	ExecTime   float64                `json:"ExecTime"`
	Headers    map[string]string      `json:"Headers"`
	Messages   []string               `json:"Messages"`
	Database   []Database             `json:"Database"`
}

// Database ...
type Database struct {
	Query     string  `json:"Query"`
	ExecTime  float64 `json:"ExecTime"`
	Status    string  `json:"Status"`
	RespError string  `json:"RespError"`
}
