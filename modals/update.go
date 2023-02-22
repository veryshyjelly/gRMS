package modals

type Update struct {
	// ID is the update id unique to user
	ID uint64 `json:"id"`
	// Message is the new incoming message
	Message *Message `json:"message"`
	// EditedMessage new version of message that was already sent
	EditedMessage *Message `json:"edited_message"`
	// ChatMember when status of a chat member is updated (for ex: permissions changed)
	ChatMember ChatMemberUpdated `json:"chat_member"`
	// ChatJoinRequest a request to join the chat has been sent
	ChatJoinRequest ChatJoinRequest `json:"chat_join_request"`
}

func NewUpdate(id uint64, mess *Message) *Update {
	return &Update{
		ID:      id,
		Message: mess,
	}
}