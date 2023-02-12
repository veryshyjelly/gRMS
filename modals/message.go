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
	From *User `json:"from"`
	// Chat in which Message is send
	Chat *Chat `json:"chat"`
	// Corresponds to User who originally send the Message
	ForwardedFrom *User `json:"forwarded_from,omitempty"`
	// Chat in which this Message was originally sent
	ForwardedFromChat *Chat `json:"forwarded_from_chat,omitempty"`
	// ReplyToMessage is the Message to which this Message is replied
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// EditDate date of last edit
	UpdatedAt *time.Time `json:"edit_date,omitempty"`
	// Text is the text message
	Text *string `json:"text,omitempty"`
	// Animation is the animated message (eg: gif)
	Animation *Animation `json:"Animation,omitempty"`
	// Audio is the audio message (eg: mp3)
	Audio *Audio `json:"audio,omitempty"`
	// Document is the document message (eg: pdf)
	Document *Document `json:"document,omitempty"`
	// Photo is the photo message (eg: jpg)
	Photo *Photo `json:"photo,omitempty"`
	// Sticker is the sticker message
	Sticker *Sticker `json:"sticker,omitempty"`
	// Video is the video message (eg: mp4, mkv)
	Video *Video `json:"video,omitempty"`
	// Caption is the caption in a media message
	Caption *string `json:"caption,omitempty"`
	// NewChatMembers list of new users joined in the chat
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	// LeftChatMember member who left the chat
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// NewChatTitle is the updated chat title
	NewChatTitle *string `json:"new_chat_title,omitempty"`
	// NewChatPhoto is the updated chat photo
	NewChatPhoto *Photo `json:"new_chat_photo,omitempty"`
	// DeleteChatPhoto is a service message, true when photo is deleted
	DeleteChatPhoto *bool `json:"delete_chat_photo,omitempty"`
	// GroupChatCreated is a service message, true when new group is created
	GroupChatCreated *bool `json:"group_chat_created,omitempty"`
	// VideoChatStarted is service message, true when video chat is started
	VideoChatStarted *bool `json:"video_chat_started,omitempty"`
	// VideoChatEnded is service message, true when video chat is ended
	VideoChatEnded *bool `json:"video_chat_ended,omitempty"`
}

func (m *Message) GetID() uint64 {
	return m.ID
}