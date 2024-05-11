package entity

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	DeletedAt   string   `json:"deleted_at"`
	Score       int64    `json:"score"`
}

type ProductSearch struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
}

type ProductCreateRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
}

type ProductCreateBleve struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
}

func InsertProductReqToCreate(req Product) *ProductCreateBleve {
	return &ProductCreateBleve{
		Name:        req.Name,
		Description: req.Description,
		Photos:      req.Photos,
		Link:        req.Link,
	}
}
