package modals

type Update struct {
	// ID is the update id unique to user
	ID uint64 `json:"id"`
	// Message is the new incoming message
	Message *Message `json:"message"`
	// EditedMessage new version of message that was already sent
	EditedMessage *Message `json:"edited_message"`
	// NewChatCreated a new chat has been created
	NewChatCreated *Chat `json:"new_chat_created"`
	// Error is the error message
	Error error `json:"error"`
}

func NewUpdate(id uint64, mess *Message) *Update {
	return &Update{
		ID:      id,
		Message: mess,
	}
}