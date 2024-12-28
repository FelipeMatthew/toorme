package utils

import "golang.org/x/crypto/bcrypt"

// TODO: SEE IF REALLY NEEDS TO HASH CZ I HAD ADDED AN EXTENSION IN POSTGRES TO BCRYPT
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
