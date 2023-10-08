// Package hash contains HashPass function, that hashes input password.
package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

// HashPass hashes input password.
func HashPass(password []byte) string {
	hasher := sha1.New()
	hasher.Write(password)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
