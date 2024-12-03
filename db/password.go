package db

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols   = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numbers   = "0123456789"
)

func generatePassword(length int, useNum bool, useSymbol bool) (string, error) {
	charset := alphabets

	if useNum {
		charset += numbers
	}

	if useSymbol {
		charset += symbols
	}

	password := make([]byte, length)

	for idx := range password {
		l := len(charset)
		charIdx, err := rand.Int(rand.Reader, big.NewInt(int64(l)))
		if err != nil {
			return "", fmt.Errorf("error generating random index: %v", err)
		}
		password[idx] = charset[charIdx.Int64()]
	}

	return string(password), nil
}
