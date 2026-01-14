package mailer

import (
	"bytes"
	"errors"
	"html/template"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	apiKey    string
	fromEmail string
}

func NewMailTrapClient(apikey, fromEmail string) (*mailtrapClient, error) {
	if apikey == "" {
		return &mailtrapClient{}, errors.New("api key is required")
	}

	return &mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apikey,
	}, nil
}

func (m *mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {

	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	// Sandbox
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())

	// Set email body
	message.AddAlternative("text/html", body.String())

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return -1, err
	} else {
		return 200, nil
	}

	// for i := 0; i < maxRetries; i++ {
	// 	res, err := m.client.Send(message)
	// 	if err != nil {
	// 		log.Printf("Failed to send mail to %v, attempt %d of %d", email, i+1, maxRetries)
	// 		log.Printf("Error: %v", maxRetries)

	// 		// expontential backoff
	// 		time.Sleep(time.Second * time.Duration(i+1))
	// 		continue
	// 	}

	// 	log.Printf("Email sent with status code %v", res.StatusCode)
	// 	return nil
	// }
	// return fmt.Errorf("failed to send email after %d attempts", maxRetries)

}
