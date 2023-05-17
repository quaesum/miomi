package entity

type Animal struct {
	ID          int64
	Age         int8
	Name        string
	Sex         int64
	Type        string
	Description string
	Sterilized  bool
	Vaccinated  bool
	Shelter     string
	Photo       string
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
