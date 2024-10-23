package templates

import (
	"bytes"
	"embed"
	"io"
	"log"
	"text/template"
)

// TemplateName is a type for the name of a template.
type TemplateName string

var (
	// RegistrationTemplate is the name of the registration template.
	RegistrationTemplate TemplateName = "registration"
	// PasswordResetTemplate is the name of the password reset template.
	PasswordResetTemplate TemplateName = "password-reset"
)

//go:embed dist/*
var static embed.FS

// GetTemplate reads the content of a file at the specified path and applies the provided values to a template.
// It returns the resulting string and any error encountered during the process.
func GetTemplate(name TemplateName, values map[string]string) (string, error) {
	log.Printf("values: %+v", values)

	file, err := static.Open("dist/" + string(name) + ".html")
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("content").Parse(string(content))
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, values)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
