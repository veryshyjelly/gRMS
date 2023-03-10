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
	// User is the user data requested
	User *User `json:"user"`
	// Chat is the chat data requested
	Chat *Chat `json:"chat"`
	// Error is the error message
	Error string `json:"error"`
}

func MessageUpdate(m *Message) *Update {
	return &Update{Message: m}
}

func ErrorUpdate(err string) *Update {
	return &Update{Error: err}
}

func NewChatUpdate(chat *Chat) *Update {
	return &Update{NewChatCreated: chat}
}

func ChatUpdate(chat *Chat) *Update {
	return &Update{Chat: chat}
}

func UserUpdate(user *User) *Update {
	return &Update{User: user}
}
