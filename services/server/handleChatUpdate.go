package server

import (
	"fmt"
	"gRMS/modals"
)

type ChatQuery struct {
	ChatId         uint64 `json:"chat_id"`
	NewTitle       string `json:"title"`
	NewDescription string `json:"description"`
}

// HandleNewTitle handles the query to change the chat title
func (sr *dvs) HandleNewTitle(c Client, query *ChatQuery) {
	// first check if the chat exists
	chat, err := sr.getChat(query.ChatId)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("chat not found"))
		return
	}
	// then check if the user has the permission to change the chat title
	if _, ok := chat.GetAdmins()[c.GetUserID()]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not admin of this chat")
		return
	}
	// now query the database to update the chat title
	err = sr.Dbs.UpdateChat(&modals.Chat{ID: query.ChatId, Title: query.NewTitle})
	if err != nil { // if any error occurs then notify the user accordingly
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error ocurred while updating chat"))
	} else { // notify all the users in the chat about the change of chat title
		sr.SendMess(&modals.Message{Chat: query.ChatId, NewChatTitle: &query.NewTitle})
	}
}