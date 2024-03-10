package utilities

import (
	"crypto/aes"
	"crypto/cipher"
)

func DecryptFile(cipherText []byte, key []byte) ([]byte, error) {
	// Create block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Deattach nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]

	cipherText = cipherText[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
