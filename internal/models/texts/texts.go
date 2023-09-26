package texts

import "time"

type Texts struct {
	UserID    string
	Name      string
	Text      []byte
	UpdatedAt time.Time
	Meta      map[string]string
}
