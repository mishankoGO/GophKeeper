package generator

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Key with public and private keys.
type Key struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// New function creates new Key instance.
func New() (Key, error) {
	var k Key

	// generate new private key with the length of 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return k, err
	}

	// retrieve the public key
	k.publicKey = &privateKey.PublicKey
	k.privateKey = privateKey

	return k, nil
}

// PublicKeyToPemString method returns pem encoding of public key.
func (k Key) PublicKeyToPemString() string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(k.publicKey),
			},
		),
	)
}

// PrivateKeyToPemString method returns pem encoding of private key.
func (k Key) PrivateKeyToPemString() string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(k.privateKey),
			},
		),
	)
}

// SaveKeys saves keys to certs folder.
func (k Key) SaveKeys() error {
	privateKey := []byte(k.PrivateKeyToPemString())
	err := os.WriteFile("certs/private.key", privateKey, 0644)
	if err != nil {
		return fmt.Errorf("error saving private key: %v", err)
	}

	publicKey := []byte(k.PublicKeyToPemString())
	err = os.WriteFile("certs/public.key", publicKey, 0644)
	if err != nil {
		return fmt.Errorf("error saving public key: %v", err)
	}
	return nil
}
