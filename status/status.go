package status

import (
	"runtime"
	"time"
)

// Provider provides status information about the go application
type Provider struct {
	start time.Time
}

// New creates a status.Provider
func New() *Provider {
	return &Provider{
		start: time.Now(),
	}
}

func (s *Provider) GetStatusInformation() map[string]interface{} {
	return map[string]interface{}{
		"started":  s.start.Format(time.RFC3339),
		"uptime":   time.Now().Sub(s.start).Minutes(),
		"rVersion": runtime.Version(),
		"now":      time.Now().Format(time.RFC3339),
	}
}
