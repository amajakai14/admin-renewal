package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pwd := []byte(password)
	hash, err:= bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	return string(hash), err
}

func MatchPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
