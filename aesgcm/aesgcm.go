package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"time"
)

type aesgcm struct {
	Key        []byte
	Length     int
	Charset    string
	SeededRand *rand.Rand
}

func NewAESGCM(key []byte) IAESGCM {
	return &aesgcm{
		Key:        key,
		Length:     12,
		Charset:    "abcdefghijklmnopqrstuvwxyz0123456789",
		SeededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type IAESGCM interface {
	GenerateDynamicIV() []byte
	EncryptWithDynamicIV(src string, iv []byte) (string, error)
	DecryptWithDynamicIV(src string, iv []byte) (string, error)
}

func (d *aesgcm) GenerateDynamicIV() []byte {
	b := make([]byte, d.Length)
	for i := range b {
		b[i] = d.Charset[d.SeededRand.Intn(len(d.Charset))]
	}
	return b
}

func (d *aesgcm) EncryptWithDynamicIV(src string, iv []byte) (string, error) {
	plaintext := []byte(src)

	block, err := aes.NewCipher(d.Key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, iv, plaintext, nil)
	result64 := base64.StdEncoding.EncodeToString(ciphertext)

	return result64, nil
}

func (d *aesgcm) DecryptWithDynamicIV(src string, iv []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(d.Key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
