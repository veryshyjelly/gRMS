package server

import (
	"chat-app/modals"
	"log"
)

func (dvs *DvService) SendMess(msg *modals.Message) {
	// This function sends the message where it needs to go
	if channel, ok := dvs.Channels[msg.Chat]; ok {
		channel.Mess <- msg
	} else {
		log.Println("error channel not found")
	}
}

func (dvs *DvService) SendChat(chat *modals.Chat) {
	// This function sends the update that the
	// chat has been created to all the users
	if c, ok := DVSr.ActiveChannels()[chat.ID]; ok {
		for client := range c.Users {
			client.Updates() <- &modals.Update{
				NewChatCreated: chat,
			}
		}
	}
}