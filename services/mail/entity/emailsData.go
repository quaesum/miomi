package entity

type EmailChangeData struct {
	Token string
}

type RegistrationConfirmData struct {
	Token string
}

type RegistrationSuccessData struct {
	Link string
}
