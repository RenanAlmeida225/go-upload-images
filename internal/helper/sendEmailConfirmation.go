package helper

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

type emailInfo struct {
	Username string
	URL      string
	Token    string
}

func parseEmail(emailInfo emailInfo) (string, error) {
	t, err := template.ParseFiles("assets/confimation.html")
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, emailInfo); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendEmailConfimation(username, email, token string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{email}
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port

	subject := "Subject: Confirm your email!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body, err := parseEmail(emailInfo{Username: username, URL: os.Getenv("URL"), Token: token})
	if err != nil {
		return err
	}
	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, host)
	if err := smtp.SendMail(address, auth, from, to, message); err != nil {
		return err
	}
	return nil
}
