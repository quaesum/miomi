package entity

type Animal struct {
	ID          int64    `json:"id"`
	Age         int8     `json:"age"`
	Name        string   `json:"name"`
	Sex         int64    `json:"sex"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Sterilized  bool     `json:"sterilized"`
	Vaccinated  bool     `json:"vaccinated"`
	OnRainbow   bool     `json:"on_rainbow"`
	OnHappiness bool     `json:"on_happiness"`
	Shelter     string   `json:"shelter,omitempty"`
	Photos      []string `json:"photos,omitempty"`
}

type AnimalCreateRequest struct {
	Age         int64
	Name        string
	Sex         int64
	Type        string
	Description string
	Sterilized  bool
	Vaccinated  bool
	Shelter     int64
}

type Shelter struct {
	ID          int64
	Logo        string
	Name        string
	Location    string
	Description string
}
type ShelterCreateRequest struct {
	Logo        string
	Name        string
	Location    string
	Description string
}

type News struct {
	ID          int64
	Label       int64
	Description string
	Photo       string
}
