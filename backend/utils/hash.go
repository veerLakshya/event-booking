package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(newPass), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
