package log_passes

import "time"

type LogPasses struct {
	UserID    string
	Name      string
	Login     []byte
	Password  []byte
	UpdatedAt time.Time
	Meta      map[string]string
}
