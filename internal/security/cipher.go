// Package security offers functionality to cipher and decipher data.
// It also has jwt manager.
package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// generateRandom generates random sequence.
func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error creating random sequence: %w", err)
	}
	return b, nil
}

// Security collects data necessary to cipher and decipher data.
type Security struct {
	key    [32]byte // secret key
	nonce  []byte
	aesgcm cipher.AEAD // block cipher mode
}

// NewSecurity function creates Security instance.
func NewSecurity(keyPhrase string) (*Security, error) {
	// hash key phrase
	key := sha256.Sum256([]byte(keyPhrase))

	// create new block
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("error creating new cipher block: %w", err)
	}

	// create new gcm
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, fmt.Errorf("error creating new gcm: %w", err)
	}

	// generate nonce
	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		return nil, fmt.Errorf("error generating random sequence: %w", err)
	}

	return &Security{
		key:    key,
		nonce:  nonce,
		aesgcm: aesgcm,
	}, nil
}

// EncryptData method ciphers input data.
func (s *Security) EncryptData(buf bytes.Buffer) []byte {
	fmt.Println("encrypt nonce: ", s.nonce, s.key)
	encData := s.aesgcm.Seal(nil, s.nonce, buf.Bytes(), nil)
	return encData
}

// DecryptData method deciphers input data.
func (s *Security) DecryptData(encData []byte) ([]byte, error) {
	fmt.Println("decrypt nonce: ", s.nonce, s.key)
	decData, err := s.aesgcm.Open(nil, s.nonce, encData, nil) // расшифровываем
	if err != nil {
		return nil, fmt.Errorf("error decrypting data: %w", err)
	}

	return decData, nil
}
