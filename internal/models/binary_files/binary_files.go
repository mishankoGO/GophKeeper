package binary_files

import "time"

type Files struct {
	UserID    string
	Name      string
	HashFile  string
	UpdatedAt time.Time
	Meta      map[string]string
}
