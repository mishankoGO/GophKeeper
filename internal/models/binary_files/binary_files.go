// Package binary_files contains Files model.
package binary_files

import "time"

// Files model.
type Files struct {
	UserID    string            `json:"user_id"`    // user id
	Name      string            `json:"name"`       // name
	File      []byte            `json:"file"`       // binary file
	Extension []byte            `json:"extension"`  // file extension
	UpdatedAt time.Time         `json:"updated_at"` // creation or update time
	Meta      map[string]string `json:"meta"`       // metadata
}
