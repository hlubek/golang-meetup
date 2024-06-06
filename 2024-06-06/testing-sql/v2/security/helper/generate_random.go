package helper

import (
	"crypto/rand"
	"errors"
)

var ErrInvalidLength = errors.New("length must be greater zero")

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	// It applies that err == nil only if we read len(bytes) bytes.
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GenerateRandomString returns a securely generated random string (base 64).
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	if n <= 0 {
		return "", ErrInvalidLength
	}

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
