package random

import (
	"crypto/rand"
	"encoding/hex"
)

// Hex generates a random hex string
func Hex(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Bytes generates a random bytes string
func Bytes(length int) ([]byte, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
