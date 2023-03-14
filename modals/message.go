package modals

import (
	"time"
)

type Message struct {
	// ID unique to Chat
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Date of Message sent
	CreatedAt time.Time `json:"timestamp"`
	// From is the Message sender
	From uint64 `json:"from"`
	// Chat in which Message is send
	Chat uint64 `json:"chat"`
	// Corresponds to User who originally send the Message
	ForwardedFrom uint64 `json:"forwarded_from,omitempty"`
	// Chat in which this Message was originally sent
	ForwardedFromChat uint64 `json:"forwarded_from_chat,omitempty"`
	// ReplyToMessage is the Message to which this Message is replied
	ReplyToMessage uint64 `json:"reply_to_message,omitempty"`
	// EditDate date of last edit
	UpdatedAt time.Time `json:"edit_date,omitempty"`
	// Text is the text message
	Text *string `json:"text,omitempty"`
	// Animation is the animated message (eg: gif)
	Animation uint64 `json:"Animation,omitempty"`
	// Audio is the audio message (eg: mp3)
	Audio uint64 `json:"audio,omitempty"`
	// Document is the document message (eg: pdf)
	Document uint64 `json:"document,omitempty"`
	// Photo is the photo message (eg: jpg)
	Photo uint64 `json:"photo,omitempty"`
	// Sticker is the sticker message
	Sticker uint64 `json:"sticker,omitempty"`
	// Video is the video message (eg: mp4, mkv)
	Video uint64 `json:"video,omitempty"`
	// Caption is the caption in a media message
	Caption *string `json:"caption,omitempty"`
	// NewChatMember is the new member added to the chat
	NewChatMember uint64 `json:"new_chat_member,omitempty"`
	// LeftChatMember member who left the chat
	LeftChatMember uint64 `json:"left_chat_member,omitempty"`
	// NewChatTitle is the updated chat title
	NewChatTitle *string `json:"new_chat_title,omitempty"`
	// NewChatPhoto is the updated chat photo
	NewChatPhoto uint64 `json:"new_chat_photo,omitempty"`
	// DeleteChatPhoto is a service message, true when photo is deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// GroupChatCreated is a service message, true when new group is created
	NewChatCreated bool `json:"new_chat_created,omitempty"`
}

func (m *Message) GetID() uint64 {
	return m.ID
}