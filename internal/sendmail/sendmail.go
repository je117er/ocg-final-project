package sendmail

import (
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	apiKey           = "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"
	fromName         = "Admin"
	fromAddress      = "thanbv1510@gmail.com"
	subject          = "Injection schedule tomorrow"
	plainTextContent = "Thank you so much <3"
	htmlContent      = `<p>Hello,</p>
					<p>You have the covid-19 vaccine scheduled for tomorrow. Please come to the clinic according to your registered schedule. <p>
					<p>thanks again,</p>
					<strong>Admin</strong>`
)

var logger = utils.SugarLog()

func SendEmailThankYou(toUsername, toEmail string) (int, error) {
	from := mail.NewEmail(fromName, fromAddress)
	to := mail.NewEmail(toUsername, toEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return -1, err

	}
	logger.Info(response)

	return response.StatusCode, nil
}
