package entity

const ConfirmEmailType = "EConf"
const RecoverEmailType = "ERecover"
const NormalEmailType = "ENorm"

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type EmailChangeData struct {
	Token string
}

type RegistrationConfirmData struct {
	Token string
}
type MailQueMessage struct {
	Type string
	To   string
	CommonData
	Meta string
}

type CommonData struct {
	Token string
}
