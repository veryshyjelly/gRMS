package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	// PhoneNumber is the phone number of the contact
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// ID uniquely identifies this contact
	ID uint64 `json:"id" gorm:"primaryKey"`
}

type ContactDB struct {
	Contact
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewContact function to create new contact entry
func NewContact(db *gorm.DB, phoneNumber, firstName, lastName string) *Contact {
	contact := ContactDB{
		Contact: Contact{
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
			LastName:    lastName,
		},
	}

	db.Create(&contact)

	return &contact.Contact
}

// FindContact used to find contact by id
func FindContact(db *gorm.DB, contactID uint64) (*ContactDB, error) {
	contact := ContactDB{Contact: Contact{ID: contactID}}
	db.First(contact)

	if contact.PhoneNumber == "" {
		return nil, fmt.Errorf("invalid contact id: %v", contact)
	}

	return &contact, nil
}