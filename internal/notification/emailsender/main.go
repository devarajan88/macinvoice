package main

import (
	"context"
	_ "embed"
	mailgun "github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

//go:embed emailtemplate.html
var html string

var (
	mailgunEmailDomain   = "sandbox9004dc52269b46948d9b08bc2f6166c4.mailgun.org"
	mailgunPrivateAPIKey = "8e09ca046b70a10a51ca16ab63d11431-70c38fed-fca8105c"
)

func main() {
	from := "test-user@reputatio.de"
	to := "devarajangct@hotmail.com"

	mg := mailgun.NewMailgun(mailgunEmailDomain, mailgunPrivateAPIKey)

	message := mg.NewMessage(from, "TEST MAIL", "", to)
	message.SetHtml(html)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//Send the message with a 10 seconds timeout
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		log.Println("error sending an email: ", err)
	}
	log.Println("Email sent via mailgun successfully")
}
