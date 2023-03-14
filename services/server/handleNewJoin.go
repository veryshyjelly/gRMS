package server

import (
	"fmt"
	"gRMS/modals"
	"log"
)

type UserQuery struct {
	ChatID uint64   `json:"chat_id"`
	Users  []string `json:"users"`
}

// HandleAddToChat adds a user to a chat and sends the chat to the user
func (sr *dvs) HandleAddToChat(c Client, query *UserQuery) {
	chat, err := sr.getChat(query.ChatID)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding chat: %v", err))
		return
	}

	if _, ok := chat.GetAdmins()[c.GetUserID()]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not an admin of this chat")
		return
	}

	for _, user := range query.Users {
		if u, err := sr.findUser(user); err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("user not found: %v", user))
		} else {
			if _, err = sr.addMember(chat.ID, u.ID); err != nil {
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error adding user to chat: %v", err))
			} else {
				sr.SendMess(&modals.Message{NewChatMember: u.ID, Chat: chat.ID})
				if p, ok := sr.ActiveUsers()[u.ID]; ok {
					p.ChatJoin() <- chat.ID
					if channel, ok := sr.ActiveChannels()[chat.ID]; ok {
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