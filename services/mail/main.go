package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"html/template"
	"log"
	entity "madmax/services/mail/entity"
	"net/smtp"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

func main() {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	consumer, err := rabbitmq.NewConsumer(
		conn,
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("meme.emails"),
		rabbitmq.WithConsumerOptionsExchangeName("emails"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		log.Printf("consumed: %v", string(d.Body))
		var mqm entity.MailQueMessage
		err = json.Unmarshal(d.Body, &mqm)
		if err != nil {

		}
		var ok bool
		for {
			ok = SendHTML(mqm)
			if ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
}

func SendHTML(mqm entity.MailQueMessage) bool {
	//password := "xsmtpsib-787d74d7be018e77f20fe1e47346a3f3850fc882e07d6b1672f04504d190da63-R5C9hL6rUWYN8Zv7"
	//from := "help.miomiby@gmail.com"
	//login := "74a5d6001@smtp-brevo.com"
	//host := "smtp-relay.brevo.com"
	//port := "587"
	//subject := "MioMi"
	host := "smtp.yandex.ru"
	password := "ecxgpvbvslizweig"
	from := "noreply-miomi@yandex.ru"
	login := "noreply-miomi"
	port := "587"
	subject := "MioMi"

	var htmlContent string
	var err error
	switch mqm.Type {
	case entity.ConfirmEmailType:
		ecd := entity.EmailChangeData{
			Token: mqm.Token,
		}
		htmlContent, err = parseTemplate("services/mail/templates/mail_verification.html", ecd)
		if err != nil {
			return false
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
	log.Println("Successful, the mail was sent!")

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
