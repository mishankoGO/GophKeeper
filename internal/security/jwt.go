// Package security offers functionality to cipher and decipher data.
// It also has jwt manager.
package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// UserClaims collects login and claims.
type UserClaims struct {
	jwt.StandardClaims        // jwt claims
	Login              string `json:"login"` // user login
}

// JWTManager contains secret key and token duration.
type JWTManager struct {
	secretKey     string        // jwt secret key
	tokenDuration time.Duration // token duration
}

// NewJWTManager function creates new jwt manager instance.
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

// Generate method generates token.
func (manager *JWTManager) Generate(user *users.User) (string, error) {
	// set claims
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Login: user.Login,
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(manager.secretKey))
}

// Verify method validates input token.
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	// parse token
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(manager.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// retrieve claims
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
