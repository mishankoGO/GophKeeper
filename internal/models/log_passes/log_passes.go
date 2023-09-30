package log_passes

import "time"

type LogPasses struct {
	UserID    string            `json:"user_id"`
	Name      string            `json:"name"`
	Login     []byte            `json:"login"`
	Password  []byte            `json:"password"`
	UpdatedAt time.Time         `json:"updated_at"`
	Meta      map[string]string `json:"meta"`
}
