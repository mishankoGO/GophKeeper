package users

import "time"

type User struct {
	UserID    string
	Login     string
	CreatedAt time.Time
}
