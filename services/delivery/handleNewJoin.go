package delivery

import (
	"chat-app/modals"
	"chat-app/services/db"
	"fmt"
	"log"
)

type AddUserQuery struct {
	ChatID uint64 `json:"chat_id"`
	UserID uint64 `json:"user_id"`
}

// HandleAddToChat adds a user to a chat and sends the chat to the user
func (c *Client) HandleAddToChat(query *AddUserQuery) {
	chat, err := dbService.DBSr.GetChat(query.ChatID)
	if err != nil {
		c.Mess <- &modals.Update{Error: fmt.Sprintf("error finding chat: %v", err)}
		return
	}

	if _, ok := chat.GetAdmins()[c.User.ID]; !ok {
		c.Mess <- &modals.Update{Error: fmt.Sprintf("unauthorized attempt to add user to chat")}
		return
	}

	_, err = dbService.DBSr.AddMember(chat.ID, query.UserID)
	if err != nil {
		c.Mess <- &modals.Update{Error: fmt.Sprintf("error adding user to group %v", err)}
	}

	if p, ok := DVSr.Users[query.UserID]; ok {
		p.Join <- chat.ID
		if channel, ok := DVSr.Channels[chat.ID]; ok {
			channel.Join <- p
			p.Mess <- &modals.Update{NewChatCreated: chat}
		} else {
			log.Fatalln("channel not found")
		}
	}
}