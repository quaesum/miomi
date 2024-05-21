package entity

import (
	"github.com/wagslane/go-rabbitmq"
)

const (
	ConfirmEmailType        = "EConf"
	RecoverEmailType        = "ERecover"
	RegistrationConfirmType = "RegConf"
)

var EmailPublisher *rabbitmq.Publisher

type MailQueMessage struct {
	Type string
	To   string
	CommonData
	Meta string
}

type CommonData struct {
	Token string
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}
