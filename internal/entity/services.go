package entity

type Service struct {
	ID          int64    `json:"id"`
	VolunteerID string   `json:"volunteer_id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	DeletedAt   string   `json:"deleted_at"`
	Score       int64    `json:"score"`
}

type ServiceCreateRequest struct {
	Label       string  `json:"label"`
	Description string  `json:"description"`
	Photos      []int64 `json:"photos"`
}
