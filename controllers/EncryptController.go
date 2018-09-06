package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"inotas-back/models"
)

type EncryptController struct {
	pass []byte
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (controller EncryptController) Encrypt(data []byte) (ciphertext string, error models.Error) {

	block, _ := aes.NewCipher([]byte(createHash(string(controller.pass[:]))))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		error = models.Error{
			Message:"Internal error",
			Code:505,
		}
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		CheckFail(err)
	}
	b := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = hex.EncodeToString(b)
	return
}

func (controller EncryptController) Decrypt(hexData string) (string, error) {
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return "", err
	}
	key := []byte(createHash(string(controller.pass[:])))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext[:]), err
}
