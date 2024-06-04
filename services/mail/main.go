package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	na "github.com/nats-io/nats.go"
	"html/template"
	"log"
	environment "madmax/sdk/environment/config"
	"madmax/sdk/nats"
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

	v, err := environment.LoadConfig(".env")
	if err != nil {
		log.Printf("Unable to load config %v", err)
		return
	}

	var settings entity.Settings
	err = v.Unmarshal(&settings)
	if err != nil {
		log.Printf("Unable to unmarshal settings %v", err)
		return
	}

	log.Println("settings", settings.NatsUrl)

	nt, err := nats.NewClient(settings.NatsUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = nt.Subscribe("mail_topic", func(msg *na.Msg) {
		re := nats.Request{}
		err := json.Unmarshal(msg.Data, &re)
		if err != nil {
			log.Printf("nats start error %v", err)
		}
		var ok bool
		for {
			ok = SendHTML(re)
			if ok {
				break
			}
			time.Sleep(time.Second * 2)
		}

		if err != nil {
			errorResp := nats.Response{Code: 400, Error: err.Error()}
			b, err := json.Marshal(errorResp)
			if err != nil {
				log.Printf("nats resp error %v", err)
			}
			err = msg.Respond(b)
			if err != nil {
				log.Printf("nats resp error %v", err)
			}
		}
		okResp := nats.Response{Code: 200}
		b, err := json.Marshal(okResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	})
	if err != nil {
		log.Printf("nats start error %v", err)
		return
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

	log.Println("awaiting signal")
	<-done
	log.Println("stopping consumer")

	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	nt.Close()

	log.Println("api has been shutdown")
}

func SendHTML(mqm nats.Request) bool {
	host := "smtp.yandex.ru"
	password := "ecxgpvbvslizweig"
	from := "noreply-miomi@yandex.ru"
	login := "noreply-miomi"
	port := "587"
	subject := "MioMi"

	var htmlContent string
	var err error
	switch mqm.Code {
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
