package server

import (
	"fmt"
	"gRMS/modals"
	dbService "gRMS/services/db"
	"log"
)

type UserQuery struct {
	ChatID uint64   `json:"chat_id"`
	Users  []string `json:"users"`
}

// HandleAddToChat adds a user to a chat and sends the chat to the user
func HandleAddToChat(c Client, query *UserQuery) {
	chat, err := dbService.DBSr.GetChat(query.ChatID)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding chat: %v", err))
		return
	}

	if _, ok := chat.GetAdmins()[c.GetUserID()]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not an admin of this chat")
		return
	}

	for _, user := range query.Users {
		if u, err := dbService.DBSr.FindUser(user); err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("user not found: %v", user))
		} else {
			if _, err = dbService.DBSr.AddMember(chat.ID, u.ID); err != nil {
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error adding user to chat: %v", err))
			} else {
				if p, ok := DVSr.ActiveUsers()[u.ID]; ok {
					p.ChatJoin() <- chat.ID
					if channel, ok := DVSr.ActiveChannels()[chat.ID]; ok {
						channel.UserJoin() <- p
						p.Updates() <- modals.NewChatUpdate(chat)
					} else {
						log.Fatalln("channel not found")
					}
				}
			}
		}
	}

}
