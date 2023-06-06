package entity

type Shelter struct {
	ID          int64
	Name        string
	Description string
	Logo        string
	Address     string
	Phone       string
	Email       string `json:"email"`
}
type ShelterCreateRequest struct {
	Name        string
	Description string
	Logo        string
	Address     string
	Phone       string
	Email       string `json:"email"`
}
