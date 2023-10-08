// Package texts contains Texts model.
package texts

import "time"

// Texts model.
type Texts struct {
	UserID    string            `json:"user_id"`    // user id
	Name      string            `json:"name"`       // text name
	Text      []byte            `json:"text"`       // text
	UpdatedAt time.Time         `json:"updated_at"` // creation or update time
	Meta      map[string]string `json:"meta"`       // metadata
}
