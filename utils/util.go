package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func Encrypt(code string) string {
	key := []byte("MovieWorkNow2022")
	plaintext := []byte(code)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("MWN")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	ciphertextStr := string(ciphertext[:])
	return ciphertextStr
}

func decrypt(code string) string {
	key := []byte("MovieWorkNow2022")
	ciphertext, _ := hex.DecodeString(code)
	nonce := []byte("MWN")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	plaintextStr := string(plaintext[:])
	return plaintextStr
}
