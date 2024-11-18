package mailer

import (
	"bytes"
	"embed"
	"log"
	"text/template"
	"time"

	"github.com/wneessen/go-mail"
)

// indicate to Go that we want to store the contents of the ./templates directory in the templateFS embeded file system variable
//
//go:embed "templates"
var templateFS embed.FS // TODO learn more about goembed

type Mailer struct {
	client *mail.Client
}

func New(host string, port int, username, password string) Mailer {
	client, err := mail.NewClient(
		host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	return Mailer{
		client: client,
	}
}

func (m Mailer) Send(sender, recipient, templateFile string, data any) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(plainBody, "plainBody", data); err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(htmlBody, "htmlBody", data); err != nil {
		return err
	}

	message := mail.NewMsg()
	if err := message.From(sender); err != nil {
		return err
	}

	if err := message.To(recipient); err != nil {
		return err
	}

	message.Subject(subject.String())
	message.SetBodyString(mail.TypeTextPlain, plainBody.String())
	message.AddAlternativeString(mail.TypeTextHTML, htmlBody.String())

	if err := m.client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
