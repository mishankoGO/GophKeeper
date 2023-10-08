// Package users contains User model.
package users

import "time"

// User model.
type User struct {
	UserID    string    `json:"user_id"`    // user id
	Login     string    `json:"login"`      // login
	CreatedAt time.Time `json:"created_at"` // creation time
}
