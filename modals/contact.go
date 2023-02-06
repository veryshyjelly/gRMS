package modals

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      uint64 `json:"user_id"`
}

func NewContact() *Contact {
	return &Contact{}
}