package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

type hashMessage struct {
	Pepper     string
	Length     int
	Charset    string
	SeededRand *rand.Rand
}

func NewHashMessage(pepper string) IHashMessage {
	return &hashMessage{
		Pepper:     pepper,
		Length:     16,
		Charset:    "abcdefghijklmnopqrstuvwxyz0123456789",
		SeededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type IHashMessage interface {
	GenerateSalt() []byte
	HashSha256EncodeNoSalt(text string) string
	HashSha256EncodeWithSalt(text string, bSalt []byte) string
	HashSha256EncodeWithPepper(text string) string
}

func (hm *hashMessage) GenerateSalt() []byte {
	b := make([]byte, hm.Length)
	for i := range b {
		b[i] = hm.Charset[hm.SeededRand.Intn(len(hm.Charset))]
	}
	return b
}

func (hm *hashMessage) HashSha256EncodeNoSalt(text string) string {
	h := sha256.Sum256([]byte(text))
	return base64.StdEncoding.EncodeToString(h[:])
}

func (hm *hashMessage) HashSha256EncodeWithSalt(text string, bSalt []byte) string {
	h := sha256.Sum256([]byte(text + string(bSalt)))
	return base64.StdEncoding.EncodeToString(h[:])
}

func (hm *hashMessage) HashSha256EncodeWithPepper(text string) string {
	h := sha256.Sum256([]byte(text + hm.Pepper))
	return base64.StdEncoding.EncodeToString(h[:])
}
