// Package cards contains Cards model.
package cards

import "time"

// Cards model.
type Cards struct {
	UserID    string            `json:"user_id"`    // user id
	Name      string            `json:"name"`       // card name
	Card      []byte            `json:"card"`       // bank card
	UpdatedAt time.Time         `json:"updated_at"` // creation or update time
	Meta      map[string]string `json:"meta"`       // metadata
}
