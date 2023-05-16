package entity

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Password  string `json:"-"`
	CreatedAt int64
	Email     string
}
type UserCreateRequest struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
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
