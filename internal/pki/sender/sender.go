package sender

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Encrypt function encrypts the input plainText message.
func Encrypt(publicKeyPath string, data []byte) (string, error) {
	// read public key file
	bytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return "", fmt.Errorf("error reading public key file: %w", err)
	}

	// convert bytes to public rsa.PublicKey
	publicKey, err := convertBytesToPublicKey(bytes)
	if err != nil {
		return "", fmt.Errorf("error converting puplic key bytes to rsa.PublicKey: %w", err)
	}

	// encrypt the message using RSA-OAEP
	msgLen := len(data)

	var encryptedBytes []byte
	random := rand.Reader
	hash := sha512.New()
	step := publicKey.Size() - 2*hash.Size() - 2

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, publicKey, data[start:finish], nil)
		if err != nil {
			return "", fmt.Errorf("error encrypting new block: %w", err)
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return cipherToPemString(encryptedBytes), nil
}

func convertBytesToPublicKey(keyBytes []byte) (*rsa.PublicKey, error) {
	var err error

	block, _ := pem.Decode(keyBytes)
	blockBytes := block.Bytes

	publicKey, err := x509.ParsePKCS1PublicKey(blockBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func cipherToPemString(cipher []byte) string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "MESSAGE",
				Bytes: cipher,
			},
		),
	)
}
