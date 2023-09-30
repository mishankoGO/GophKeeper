package texts

import "time"

type Texts struct {
	UserID    string            `json:"user_id"`
	Name      string            `json:"name"`
	Text      []byte            `json:"text"`
	UpdatedAt time.Time         `json:"updated_at"`
	Meta      map[string]string `json:"meta"`
}
