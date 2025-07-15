package email

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

type EmailSenderAdapter struct {
	apiKey string
	from   string
}

func NewEmailSenderAdapter(apiKey, from string) *EmailSenderAdapter {
	return &EmailSenderAdapter{
		apiKey: apiKey,
		from:   from,
	}
}

func (s *EmailSenderAdapter) SendEmail(name, to, subject, body string) error {
	var target string
	if name != "" {
		target = fmt.Sprintf("%s <%s>", name, to)
	} else {
		target = fmt.Sprintf("%s", to)
	}

	mg := mailgun.NewMailgun(s.from, s.apiKey)
	m := mailgun.NewMessage(
		fmt.Sprintf("Lorecrafter <postmaster@%s>", s.from),
		subject,
		body,
		target,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, _, err := mg.Send(ctx, m)
	return err
}
