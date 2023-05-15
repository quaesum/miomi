package entity

type User struct {
	ID    int64
	Name  string
	Email string
}

type Animal struct {
	ID          int64
	Age         int8
	Name        string
	Sex         string
	Type        string
	Description string
	Castrated   bool
	Sterilized  bool
	Vaccinated  bool
	Shelter     string
}
