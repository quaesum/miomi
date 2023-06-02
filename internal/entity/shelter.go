package entity

type Shelter struct {
	ID          int64
	Logo        string
	Name        string
	Location    string
	Description string
}
type ShelterCreateRequest struct {
	Name        string
	Description string
	Logo        string
	Address     string
	Phone       string
	Email       string `json:"email"`
}
