package utilities

import (
	"crypto/rand"
)

func GenAESKey() ([]byte, error) {
	keyBytes := make([]byte, 32)

	_, err := rand.Read(keyBytes)
	if err != nil {
		return nil, err
	}

	return keyBytes, err
}
