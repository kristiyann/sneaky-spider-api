package logic

import (
	"github.com/kristiyann/af1-spider-web-app/util"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateEmailConfirmationCode() (string, error) {
	return util.GenerateRandomString(32)
}
