package rabbitmq

import (
	"bytes"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"html/template"
	"log"
	"madmax/internal/entity"
	"net/smtp"
	"strings"
	"unicode/utf8"
)

var EmailPublisher *rabbitmq.Publisher

func SendHTML(mqm entity.MailQueMessage) bool {
	password := "k5ZvQ5QTzs76z2HG35"
	from := "help.miomiby@gmail.com"
	login := "help.miomiby@gmail.com"
	host := "smtp-relay.gmail.com"
	port := "465"
	subject := "MioMI"

	var htmlContent string
	var err error
	switch mqm.Type {
	case entity.RecoverEmailType:
		ecd := entity.EmailChangeData{
			Token: mqm.Token,
		}
		htmlContent, err = parseTemplate("./serives/mail/template/identity_mail_verification.html", ecd)
		if err != nil {
			log.Println(err)
		}
		subject = "MioMi подтверждение почты"
	}

	message := BuildMessage(entity.Mail{
		Sender:  from,
		To:      []string{mqm.To},
		Subject: subject,
		Body:    htmlContent,
	})

	auth := smtp.PlainAuth("", login, password, host)
	err = smtp.SendMail(host+":"+port, auth, from, []string{mqm.To}, []byte(message))
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println("Successful, the mail was sent!")

	return true
}

func BuildMessage(mail entity.Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func parseTemplate(fileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	s := buffer.String()
	return s, nil
}
func StringToAsciiBytes(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		i++
	}
	return t
}

func SetupRabbitMQ() (*rabbitmq.Publisher, error) {
	// Load RabbitMQ connection details from environment variables or configuration
	// Create a new RabbitMQ connection
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ connection: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}()

	// Create a new RabbitMQ publisher
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("emails"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ publisher: %w", err)
	}

	EmailPublisher = publisher

	return publisher, nil
}
