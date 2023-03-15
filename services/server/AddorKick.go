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
	// firs check whether the chat exists or not
	chat, err := sr.Dbs.GetChat(query.ChatID)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding chat: %v", err))
		return
	}
	// now check whether the user has the right permission to add users
	if _, ok := chat.GetAdmins()[c.GetUserID()]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not an admin of this chat")
		return
	}
	// now loop over all the user ids
	for _, user := range query.Users {
		if u, err := sr.Dbs.FindUser(user); err != nil { // find the user from database
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("user not found: %v", user))
		} else {
			if _, err = sr.Dbs.AddMember(chat.ID, u.ID); err != nil { // add a relation of user and the chat
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error adding user to chat: %v", err))
			} else { // send message to all the users in the chat about new user
				sr.SendMess(&modals.Message{NewChatMember: u.ID, Chat: chat.ID})
				if p, ok := sr.ActiveUsers()[u.ID]; ok { // if the user is active then add the channel to his channels map
					p.ChatJoin() <- chat.ID
					if channel, ok := sr.ActiveChannels()[chat.ID]; ok {
						channel.UserJoin() <- p
						p.Updates() <- modals.NewChatUpdate(chat) // update the user about his new chat
					} else {
						log.Fatalln("channel not found")
					}
				}
			}
		}
	}
}

// HandleRemoveFromChat handles kicking of user
func (sr *dvs) HandleRemoveFromChat(c Client, query *UserQuery) {
	// firs check whether the chat exists or not
	chat, err := sr.Dbs.GetChat(query.ChatID)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding chat: %v", err))
		return
	}
	// now check whether the user has the right permission to add users
	if _, ok := chat.GetAdmins()[c.GetUserID()]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not an admin of this chat")
		return
	}
	// now loop over all the user ids
	for _, user := range query.Users {
		if u, err := sr.Dbs.FindUser(user); err != nil { // find the user from database
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("user not found: %v", user))
		} else {
			if err = sr.Dbs.RemoveMember(chat.ID, u.ID); err != nil { // add a relation of user and the chat
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error removing user from chat: %v", err))
			} else { // send message to all the users in the chat about new user
				sr.SendMess(&modals.Message{LeftChatMember: u.ID, Chat: chat.ID})
				if p, ok := sr.ActiveUsers()[u.ID]; ok { // if the user is active then add the channel to his channels map
					p.LeaveChat() <- chat.ID
					if channel, ok := sr.ActiveChannels()[chat.ID]; ok {
						channel.UserLeave() <- p
						//p.Updates() <- modals.NewChatUpdate(chat) // update the user about his new chat
					} else {
						log.Fatalln("channel not found")
					}
				}
			}
		}
	}
}