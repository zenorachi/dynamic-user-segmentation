package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(password string) (string, error) {
	hasher := sha1.New()

	if _, err := hasher.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum([]byte(h.salt))), nil
}
