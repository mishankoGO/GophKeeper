// Package log_passes contains LogPasses model.
package log_passes

import "time"

// LogPasses model.
type LogPasses struct {
	UserID    string            `json:"user_id"`    // user id
	Name      string            `json:"name"`       // name of log pass
	Login     []byte            `json:"login"`      // login
	Password  []byte            `json:"password"`   // password
	UpdatedAt time.Time         `json:"updated_at"` // creation or update time
	Meta      map[string]string `json:"meta"`       // metadata
}
