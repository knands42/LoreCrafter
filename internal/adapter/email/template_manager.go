package email

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templatesFS embed.FS

type TemplateManagerAdapter struct {
	templates map[string]*template.Template
}

func NewTemplateManagerAdapter() (*TemplateManagerAdapter, error) {
	tm := &TemplateManagerAdapter{
		templates: make(map[string]*template.Template),
	}

	// Load all templates
	tmpl, err := template.ParseFS(templatesFS, "templates/email_verification.html")
	if err != nil {
		return nil, err
	}
	tm.templates["email_verification"] = tmpl

	tmpl2, err := template.ParseFS(templatesFS, "templates/password_reset.html")
	if err != nil {
		return nil, err
	}
	tm.templates["password_reset"] = tmpl2

	return tm, nil
}

func (tm *TemplateManagerAdapter) Render(templateName string, data interface{}) (string, error) {
	tmpl, exists := tm.templates[templateName]
	if !exists {
		return "", nil
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
