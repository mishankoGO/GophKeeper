package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Security struct {
	key    [32]byte
	nonce  []byte
	aesgcm cipher.AEAD
}

func NewSecurity(keyPhrase string) (*Security, error) {
	key := sha256.Sum256([]byte(keyPhrase))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating new cipher block: %v", err)
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating new gcm: %v", err)
	}

	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error generating random sequence: %v", err)
	}

	return &Security{
		key:    key,
		nonce:  nonce,
		aesgcm: aesgcm,
	}, nil
}

func (s *Security) EncryptData(buf bytes.Buffer) []byte {
	encData := s.aesgcm.Seal(nil, s.nonce, buf.Bytes(), nil)
	return encData
}

func (s *Security) DecryptData(encData []byte) ([]byte, error) {
	decData, err := s.aesgcm.Open(nil, s.nonce, encData, nil) // расшифровываем
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	}

	return decData, nil
}
