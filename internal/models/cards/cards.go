package cards

import "time"

type Cards struct {
	UserID    string
	Name      string
	Card      []byte
	UpdatedAt time.Time
	Meta      map[string]string
}
