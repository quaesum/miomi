package entity

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"createdAt"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	ShelterID int64  `json:"shelter_id"`
}

type UserLogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
}

type UserResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt int64  `json:"created_at"`
	Email     string `json:"email"`
}
