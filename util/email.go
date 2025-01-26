package util

import (
	"fmt"
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/resend/resend-go/v2"
)

func SendGoodbyeEmail(toEmail string) error {
	subject := "Sneaky Spider -- Sorry to see you go :(!"
	body := `We're sorry to see you go!

	We noticed you cancelled your subscription to Sneaky Spider & would really appreciate your feedback on what we could have done better to make you stay. You can reply to this email.
	
	Thank you for being our customer!
	
	https://sneakyspiderapp.com`

	return SendEmail(toEmail, subject, body)
}

func SendPaymentFailedEmail(toEmail string) error {
	subject := "Sneaky Spider -- Payment failed."
	body := `Hello,

	Unfortunately, your latest payment attempt has failed. It is possible that your bank or credit card details are invalid.
	
	If this seems to be an error on our end, please let us know by replying to this email.
	
	https://sneakyspiderapp.com`

	return SendEmail(toEmail, subject, body)
}

func SendEmail(toEmail string, subject string, body string) error {
	// fromEmail := LoadEnvironmentVariable(constants.EnvSmtpEmail)
	// fromPassword := LoadEnvironmentVariable(constants.EnvSmtpPassword)
	// auth := smtp.PlainAuth("Sneaky Spider", fromEmail, fromPassword, "smtp.gmail.com")

	// return smtp.SendMail("smtp.gmail.com:587", auth, fromEmail, []string{toEmail}, []byte(body))
	return sendWithResend(toEmail, subject, body)
}

func sendWithResend(toEmail string, subject string, body string) error {
	apiKey := LoadEnvVar(constants.EnvResendAPIKey)
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Sneaky Spider <%s>", constants.AlertsEmail),
		To:      []string{toEmail},
		Subject: subject,
		Text:    body,
		ReplyTo: "chips4real4@gmail.com",
		// Html:    fmt.Sprintf("<span>%s<span>", body),
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return models.NewAPIError("Could not send Email with Resend: "+err.Error(), http.StatusInternalServerError)
	}

	return nil
}
