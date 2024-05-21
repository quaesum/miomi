package entity

type Shelter struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	IsVerified  bool   `json:"isVerified"`
}
type ShelterCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}

type ShelterConfirmRequest struct {
	RequestedBy int64  `json:"requestedBy"`
	ShelterID   int64  `json:"shelterID"`
	CreatedAt   string `json:"createdAt"`
}

type ShelterConfirmRequestInfo struct {
	RequestedVolunteer User    `json:"user"`
	Shelter            Shelter `json:"shelter"`
	CreatedAt          string  `json:"createdAt"`
}

type ShelterInviteRequest struct {
	RequestedBy int64 `json:"requestedBy"`
	ShelterID   int64 `json:"shelterID"`
}

type ShelterInvitation struct {
	ID        int64   `json:"id"`
	From      User    `json:"from"`
	InvitedTo Shelter `json:"invitedTo"`
}

func CompareUserToShelter(u *User) *ShelterCreateRequest {
	return &ShelterCreateRequest{
		Name:    u.FirstName + " " + u.LastName,
		Address: "",
		Phone:   u.Phone,
	}
}
