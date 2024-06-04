package entity

import (
	"github.com/nats-io/nats.go"
)

const (
	ConfirmEmailType        = "EConf"
	RecoverEmailType        = "ERecover"
	RegistrationConfirmType = "RegConf"
)

var Nats *nats.Conn

type MailQueMessage struct {
	Type string
	To   string
	CommonData
	Meta string
}

type CommonData struct {
	Token                 string
	TenderNum             string
	TenderName            string
	Link                  string
	CompanyName           string
	CompanyUNP            string
	Bidding               string
	Participators         string
	ParticipatorsListLink string
	Data                  string
	Time                  string
	ReTenderNum           string
	ReTenderName          string
	Percent               string
}
type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type Settings struct {
	NatsUrl string `json:"NATS_URL" mapstructure:"NATS_URL"`
}
