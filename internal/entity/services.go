package entity

type Service struct {
	ID          int64    `json:"id"`
	VolunteerID int64    `json:"volunteer_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	DeletedAt   string   `json:"deleted_at"`
}

type ServiceCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photos      []int64 `json:"photos"`
}

type ServiceSearch struct {
	ID          int64    `json:"id"`
	VolunteerID int64    `json:"volunteer_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"images"`
}

type ServiceCreateBleve struct {
	VolunteerID int64    `json:"volunteer_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type ServiceBleve struct {
	ID          int64    `json:"id"`
	VolunteerID float64  `json:"volunteer_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"images"`
}

func InsertServiceReqToCreate(req Service) *ServiceCreateBleve {
	return &ServiceCreateBleve{
		VolunteerID: req.VolunteerID,
		Name:        req.Name,
		Description: req.Description,
		Photos:      req.Photos,
	}
}
