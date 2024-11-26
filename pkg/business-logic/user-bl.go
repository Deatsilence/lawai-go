package businesslogic

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
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

func GenerateResetCode(ctx context.Context, email string) error {
	source := rand.NewSource(time.Now().UnixNano())
	localRNG := rand.New(source)

	code := fmt.Sprintf("%06d", localRNG.Intn(1000000))
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

}
