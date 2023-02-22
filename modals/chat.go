package modals

import "time"

type Chat struct {
	// Unique ID of the Chat
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Title of the Chat
	Title string `json:"title"`
	// Members is the list of usernames in the chat
	Members []Participant `json:"members" gorm:"foreignKey:ChatID"`
	// Admins is the list of admins in the chat
	Admins []Admin `json:"admins" gorm:"foreignKey:ChatID"`
	// DP is the display picture of the chat
	DP uint64 `json:"dp"`
	// Description is the chat description
	Description string `json:"description"`
	// InviteLink is the current active invite link
	InviteLink string `json:"invite_link"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

func (c Chat) GetAdmins() map[uint64]bool {
	admins := make(map[uint64]bool)
	for _, admin := range c.Admins {
		admins[admin.UserID] = true
	}
	return admins
}