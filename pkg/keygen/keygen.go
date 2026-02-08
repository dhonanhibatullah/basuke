package keygen

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func Encrypt(data []byte) (encrypted string) {
	hashedKey := sha256.Sum256(secretKey)
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return ""
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return ""
	}

	encByte := gcm.Seal(nonce, nonce, data, nil)
	return base64.URLEncoding.EncodeToString(encByte)
}

func Decrypt(encrypted string) (data []byte) {
	encByte, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return nil
	}

	hashedKey := sha256.Sum256(secretKey)
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil
	}

	nonceSize := gcm.NonceSize()
	if len(encByte) < nonceSize {
		return nil
	}

	nonce, ciphertext := encByte[:nonceSize], encByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil
	}

	return plaintext
}
