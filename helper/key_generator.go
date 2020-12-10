package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GenerateRSAKeyPair generate rsa key pair
func GenerateRSAKeyPair(path string) error {
	os.Mkdir(path, os.ModePerm)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePem, err := os.Create(path + "/private.pem")
	if err != nil {
		return err
	}

	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		return err
	}

	publickKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publickKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publickKeyBytes,
	}

	publicPem, err := os.Create(path + "/public.pem")
	if err != nil {
		return err
	}

	err = pem.Encode(publicPem, publickKeyBlock)
	if err != nil {
		return err
	}

	return nil
}
