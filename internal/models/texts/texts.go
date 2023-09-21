package texts

import "time"

type Texts struct {
	UserID    string
	Name      string
	HashText  string
	UpdatedAt time.Time
	Meta      map[string]string
}
