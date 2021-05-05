package user

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type UserFlags []string

func (i *UserFlags) String() string {
	return "my string representation"
}

func (i *UserFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func SendMail(email string, centerFormatted string) {
	from := mail.NewEmail("", os.Getenv("SENDGRID_FROM_EMAIL"))
	subject := "Cowin Alert!"
	to := mail.NewEmail("", email)
	plainTextContent := "Slots available"
	htmlContent := fmt.Sprintf("<pre>%s</pre>", centerFormatted)
	log.Println("sending mail to:", email, "from:", from.Address, " with env:", os.Getenv("SENDGRID_API_KEY"))
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.StatusCode)
	log.Println(resp.Body)
	log.Println(resp.Headers)
}
