package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// RememberTokenBytes stores the size (the number of bytes) of the remember tokens
const RememberTokenBytes = 32

// Bytes will generate n random bytes or return an error if it fails to do so.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String will generate a remember token string of n bytes
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken returns random string of fixed size as remember token
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}