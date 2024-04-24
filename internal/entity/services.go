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
}

type CreateServiceRequest struct {
	VolunteerID string   `json:"volunteer_id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}
