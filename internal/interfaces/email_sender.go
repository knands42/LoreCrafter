package interfaces

// EmailSender defines the interface for sending emails
type EmailSender interface {
	// SendEmail sends an email to the specified recipient
	SendEmail(name, to, subject, body string) error
}

type TemplateManager interface {
	Render(templateName string, data interface{}) (string, error)
}
