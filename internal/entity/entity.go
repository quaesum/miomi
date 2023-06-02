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
	Shelter     string   `json:"shelter"`
	Photos      []string `json:"photos"`
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

type News struct {
	ID          int64  `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Photo       string `json:"photo,omitempty"`
}
