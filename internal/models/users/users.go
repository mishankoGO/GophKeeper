package users

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"created_at"`
}
