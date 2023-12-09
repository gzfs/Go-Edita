package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword string, password string) (bool, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword), nil
}
