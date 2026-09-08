package jwtmanager

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func LoadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, errors.New("invalid PEM format")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privateKey, nil
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("private key parsing error: %w", err)
	}

	rsaPrivateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("the key is not an RSA private key")
	}

	return rsaPrivateKey, nil
}

func LoadPublicKey(fileName string) (*rsa.PublicKey, error) {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, errors.New("invalid PEM format")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("public key parsing error: %w", err)
	}
	rsaPublicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("the key is not RSA public key")
	}
	return rsaPublicKey, nil
}
