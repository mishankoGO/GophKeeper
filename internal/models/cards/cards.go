package cards

import "time"

type Cards struct {
	UserID    string            `json:"user_id"`
	Name      string            `json:"name"`
	Card      []byte            `json:"card"`
	UpdatedAt time.Time         `json:"updated_at"`
	Meta      map[string]string `json:"meta"`
}
