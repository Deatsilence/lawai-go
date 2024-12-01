package businesslogic

import (
	"golang.org/x/crypto/bcrypt"
)

type UserBL struct{}

func NewUserBL() *UserBL {
	return &UserBL{}
}

func (bl *UserBL) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
