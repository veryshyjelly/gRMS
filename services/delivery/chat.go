package delivery

import (
	"chat-app/modals"
	"chat-app/services/db"
	"fmt"
	"log"
)

type NewChatQuery struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

// HandleNewChat creates a new chat and starts a new channel to handle it
func (c *Client) HandleNewChat(chatQuery *NewChatQuery) {
	users := []uint64{c.User.ID}

	if chatQuery.Participants == nil {
		c.Mess <- &modals.Update{Error: fmt.Errorf("error while creating chat: participants is nil")}
		return
	}

	for _, username := range chatQuery.Participants {
		user, err := dbService.DBSr.FindUser(username)
		if err != nil {
			c.Mess <- &modals.Update{Error: fmt.Errorf("error while creating chat: %v", err)}
			return
		}
		users = append(users, user.ID)
	}

	chat, err := dbService.DBSr.CreateChat(users, chatQuery.Title)
	if err != nil {
		c.Mess <- &modals.Update{Error: fmt.Errorf("error while creating chat: %v", err)}
		return
	}

	c.Mess <- &modals.Update{NewChatCreated: chat}

	DVSr.mu.Lock()
	newChannel := NewChannel(chat.ID)
	DVSr.Channels[chat.ID] = newChannel
	DVSr.mu.Unlock()

	// Start the channel to handle the chat
	go newChannel.Run()
	HandleAllJoin(chat)

	DVSr.SendChat(chat)
}

// HandleAllJoin adds all users to the channel
func HandleAllJoin(chat *modals.Chat) {
	if channel, ok := DVSr.Channels[chat.ID]; ok {
		for _, user := range chat.Members {
			if client, ok := DVSr.Users[user.ID]; ok {
				client.Chats[chat.ID] = true
				channel.Join <- client
				client.Mess <- &modals.Update{NewChatCreated: chat}
			}
		}
	} else {
		log.Fatalln("channel not found")
	}
}