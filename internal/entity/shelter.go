package entity

type Shelter struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}
type ShelterCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}
