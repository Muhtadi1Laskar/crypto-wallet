package crypto

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateSeed(mnemonic, password string) []byte {
	salt := "menomic" + password
	iterations := 2048
	keyLen := 64
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)

	return pbkdf2.Key(passwordBytes, saltBytes, iterations, keyLen, sha512.New)
}