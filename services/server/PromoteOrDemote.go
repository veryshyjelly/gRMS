package server

import (
	"fmt"
	"gRMS/modals"
)

func (sr *dvs) HandlePromoteUsers(c Client, query *UserQuery) {
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
			if _, err = sr.Dbs.AddAdmin(chat.ID, u.ID); err != nil { // add a relation of user and the chat
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error promoting user in chat: %v", err))
			} else { // send message to all the users in the chat about new user
				//
			}
		}
	}
}

func (sr *dvs) HandleDemoteUsers(c Client, query *UserQuery) {
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
			if err = sr.Dbs.RemoveAdmin(chat.ID, u.ID); err != nil { // add a relation of user and the chat
				c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error demoting user in chat: %v", err))
			} else { // send message to all the users in the chat about new user
				//sr.SendMess(&modals.Message{LeftChatMember: u.ID, Chat: chat.ID})
			}
		}
	}
}