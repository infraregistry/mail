// Package mail provides functions for sending emails.
package mail

import (
	"context"
	"fmt"

	"github.com/resend/resend-go/v2"

	"github.com/nvr-ai/go-mail/templates"
)

type Recipients struct {
	To    []string
	Cc    []string
	Bcc   []string
	Reply string
}

// Send sends an email using the specified template, recipient, and data.
// It returns an error if there was a problem sending the email.
func Send(apiKey string, name templates.TemplateName, recipients Recipients, subject string, data map[string]string) error {
	template, err := templates.GetTemplate(name, data)
	if err != nil {
		return err
	}

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "support@infraregistry.com",
		ReplyTo: "support@infraregistry.com",
		To:      recipients.To,
		Html:    template,
		Subject: subject,
		Cc:      recipients.Cc,
		Bcc:     recipients.Bcc,
	}

	sent, err := client.Emails.SendWithContext(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println(sent.Id)

	return nil
}
