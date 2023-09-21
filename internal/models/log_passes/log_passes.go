package log_passes

import "time"

type LogPasses struct {
	UserID       string
	Name         string
	HashLogin    string
	HashPassword string
	UpdatedAt    time.Time
	Meta         map[string]string
}
