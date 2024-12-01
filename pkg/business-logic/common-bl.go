package businesslogic

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	models "github.com/Deatsilence/lawai-go/pkg/models/mongo"
)

type CommonBL struct{}

func NewCommonBL() *CommonBL {
	return &CommonBL{}
}

func (bl *CommonBL) GenerateResetCode(ctx context.Context, email string) *models.PasswordReset {
	source := rand.NewSource(time.Now().UnixNano())
	localRNG := rand.New(source)

	code := fmt.Sprintf("%06d", localRNG.Intn(1000000))

	passwordReset := &models.PasswordReset{
		Email:     &email,
		Code:      code,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 1), // Code expires in 1 minute
	}
	return passwordReset
}

func (bl *CommonBL) SendEmail(toEmail string, code string) error {
	fromMail := os.Getenv("FROMMAIL")
	fromMailPassword := os.Getenv("FROMMAILPASSWORD")

	auth := smtp.PlainAuth("", fromMail, fromMailPassword, "smtp.gmail.com")

	msg := "Subject: Password Reset Code\n\nHere is your password reset code: " + code

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		fromMail,
		[]string{toEmail},
		[]byte(msg),
	)

	if err != nil {
		log.Printf("Error while sending email: %v", err)
		return err
	}
	return nil
}
