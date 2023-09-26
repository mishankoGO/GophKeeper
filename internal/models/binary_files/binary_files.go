package binary_files

import "time"

type Files struct {
	UserID    string
	Name      string
	File      []byte
	UpdatedAt time.Time
	Meta      map[string]string
}
