package services

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const key string = "1234567890123456"

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(key)[:block.BlockSize()])
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, []byte(key)[:block.BlockSize()])
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
