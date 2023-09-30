package binary_files

import "time"

type Files struct {
	UserID    string            `json:"user_id"`
	Name      string            `json:"name"`
	File      []byte            `json:"file"`
	UpdatedAt time.Time         `json:"updated_at"`
	Meta      map[string]string `json:"meta"`
}
