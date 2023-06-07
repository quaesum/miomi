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
	ShelterId   int64    `json:"shelterId"`
	Photos      []string `json:"photos"`
	Address     string   `json:"adress"`
	Phone       string   `json:"phone"`
}

type AnimalCreateRequest struct {
	Age         int8    `json:"age"`
	Name        string  `json:"name"`
	Sex         int64   `json:"sex"`
	Type        int64   `json:"type"`
	Description string  `json:"description"`
	Sterilized  bool    `json:"sterilized"`
	Vaccinated  bool    `json:"vaccinated"`
	ShelterId   int64   `json:"shelterId"`
	Photos      []int64 `json:"photos"`
}

type News struct {
	ID          int64  `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Photo       string `json:"photo,omitempty"`
	CreatedAt   string `json:"created_at"`
}
