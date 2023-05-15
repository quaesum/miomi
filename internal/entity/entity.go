package entity

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Password  string
	CreatedAt int64
	Email     string
}
type UserCreateRequest struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
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
