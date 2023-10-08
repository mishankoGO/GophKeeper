// Package users contains Credential model.
package users

// Credential model.
type Credential struct {
	Login    string `json:"login"`    // login
	Password string `json:"password"` // password
}
