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
		error = models.ErrorResponse(err, 400)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		if err != nil{
			return "", models.ErrorResponse(err,400)
		}
	}

	b := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = hex.EncodeToString(b)
	return
}

func (controller EncryptController) Decrypt(hexData string) (string, models.Error) {
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return "", models.ErrorResponse(err, 500)
	}
	key := []byte(createHash(string(controller.pass[:])))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", models.ErrorResponse(err, 500)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", models.ErrorResponse(err, 500)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext[:]), models.Error{}
}
