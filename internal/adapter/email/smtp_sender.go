package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type SMTPSenderAdapter struct {
	server   string
	port     int
	username string
	password string
	from     string
}

func NewSMTPSenderAdapter(server string, port int, username, password, from string) *SMTPSenderAdapter {
	return &SMTPSenderAdapter{
		server:   server,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *SMTPSenderAdapter) SendEmail(to, subject, body string) error {
	headers := make(map[string]string)
	headers["From"] = s.from
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", s.username, s.password, s.server)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.server,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", s.server, s.port), tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to dial SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.server)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// Auth
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	// Set the sender and recipient
	if err = client.Mail(s.from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Send the email body
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return nil
}
