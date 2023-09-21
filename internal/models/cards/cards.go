package cards

import "time"

type Cards struct {
	UserID         string
	Name           string
	HashCardNumber string
	HashCardHolder string
	ExpiryDate     time.Time
	HashCVV        string
	UpdatedAt      time.Time
	Meta           map[string]string
}
