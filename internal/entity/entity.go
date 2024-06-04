package entity

import "strconv"

type Animal struct {
	ID          int64    `json:"id"`
	Age         int8     `json:"age"`
	AgeType     string   `json:"ageType"`
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
	Address     string   `json:"address"`
	Phone       string   `json:"phone"`
	Score       int64    `json:"score"`
}

type AnimalFilters struct {
	Sex        []string `json:"sex"`
	Type       []string `json:"type"`
	Sterilized bool     `json:"sterilized"`
	Vaccinated bool     `json:"vaccinated"`
	ShelterId  []int    `json:"shelterId"`
	MaxAge     float64  `json:"maxAge"`
	MinAge     float64  `json:"minAge"`
}

type AnimalTypes struct {
	ID   int64  `json:"id"`
	Type string `json:"name"`
}

func (f *AnimalFilters) IsEmpty() bool {
	return len(f.Sex) == 0 &&
		len(f.Type) == 0 &&
		!f.Sterilized &&
		!f.Vaccinated &&
		len(f.ShelterId) == 0 &&
		f.MaxAge == 0 &&
		f.MinAge == 0
}

type AnimalCreateBleve struct {
	Age         int      `json:"age"`
	AgeType     string   `json:"ageType"`
	Name        string   `json:"name"`
	Sex         string   `json:"sex"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Sterilized  bool     `json:"sterilized"`
	Vaccinated  bool     `json:"vaccinated"`
	ShelterId   string   `json:"shelterId"`
	Shelter     string   `json:"shelter"`
	Address     string   `json:"address"`
	Photos      []string `json:"photos"`
}

type AnimalBleve struct {
	ID          int64    `json:"id"`
	Age         float64  `json:"age"`
	AgeType     *string  `json:"ageType"`
	Name        string   `json:"name"`
	Sex         string   `json:"sex"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Sterilized  bool     `json:"sterilized"`
	Vaccinated  bool     `json:"vaccinated"`
	ShelterId   string   `json:"shelterId"`
	Shelter     string   `json:"shelter"`
	Address     string   `json:"address"`
	Photos      []string `json:"photos"`
}

func InserAnimalReqToCreate(req Animal) *AnimalCreateBleve {
	return &AnimalCreateBleve{
		Age:         int(req.Age),
		AgeType:     req.AgeType,
		Name:        req.Name,
		Sex:         strconv.FormatInt(req.Sex, 10),
		Type:        req.Type,
		Description: req.Description,
		Sterilized:  req.Sterilized,
		Vaccinated:  req.Vaccinated,
		ShelterId:   strconv.FormatInt(req.ShelterId, 10),
		Shelter:     req.Shelter,
		Address:     req.Address,
		Photos:      req.Photos,
	}
}

type SearchRequest struct {
	Request string        `json:"request"`
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
	Filters AnimalFilters `json:"filters"`
}

type AnimalCreateRequest struct {
	Age         int8    `json:"age"`
	AgeType     string  `json:"ageType"`
	Name        string  `json:"name"`
	Sex         int64   `json:"sex"`
	Type        int64   `json:"type"`
	Description string  `json:"description"`
	Sterilized  bool    `json:"sterilized"`
	Vaccinated  bool    `json:"vaccinated"`
	ShelterId   int64   `json:"shelterId"`
	Onrainbow   bool    `json:"onrainbow"`
	Onhappines  bool    `json:"onhappines"`
	Photos      []int64 `json:"photos"`
}

type News struct {
	ID          int64  `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Photo       string `json:"photo,omitempty"`
	CreatedAt   string `json:"created_at"`
}

type NewsCreateRequest struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Photo       int64  `json:"photo"`
}

type PhotoRequest struct {
	ID       int64  `json:"id"`
	Filename string `json:"url"`
}

type SearchAnimalsResponse struct {
	Animals []Animal `json:"animals"`
	MaxPage int8     `json:"max_page"`
}

type SearchAnimalsResponseV2 struct {
	Animals []AnimalBleve `json:"animals"`
	MaxPage int8          `json:"max_page"`
}

type SearchServicesResponse struct {
	Services []ServiceBleve `json:"services"`
	MaxPage  int8           `json:"max_page"`
}
type SearchProductsResponse struct {
	Products []ProductSearch `json:"products"`
	MaxPage  int8            `json:"max_page"`
}

type ReportCreateRequest struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

type Report struct {
	ID          int64        `json:"id"`
	Sender      UserResponse `json:"sender"`
	Label       string       `json:"label"`
	Description string       `json:"description"`
}
