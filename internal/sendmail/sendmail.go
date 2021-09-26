package sendmail

import (
	"fmt"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	apiKey           = "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"
	fromName         = "Admin"
	fromAddress      = "devnotfoundex@gmail.com"
	subject          = "Injection schedule"
	plainTextContent = "Injection schedule"
	htmlContent      = `<p>Hello %s,</p>
						<p>You have vaccination schedule on %s at %s clinic in the %s, Please come on time.</p>
						<p>Get answers to questions: %s<p>
						<p> <p>thanks again,</p>
						<strong>Admin</strong>`
)

var logger = utils.SugarLog()

func SendEmailThankYou(sendMail models.SentMail) (int, error) {
	from := mail.NewEmail(fromName, fromAddress)
	to := mail.NewEmail(sendMail.Name, sendMail.Email)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, fmt.Sprintf(htmlContent, sendMail.Name, sendMail.Date.Format("2006-01-02"), sendMail.ClinicName, sendMail.TimePeriod, "1-800-232-0233"))
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return -1, err

	}
	logger.Info(response)

	return response.StatusCode, nil
}
