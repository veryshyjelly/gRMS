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
		c.Mess <- &modals.Update{Error: fmt.Sprintf("error while creating chat: participants is nil")}
		return
	}

	for _, username := range chatQuery.Participants {
		user, err := dbService.DBSr.FindUser(username)
		if err != nil {
			c.Mess <- &modals.Update{Error: fmt.Sprintf("error while creating chat: %v", err)}
			return
		}
		users = append(users, user.ID)
	}

	chat, err := dbService.DBSr.CreateChat(users, chatQuery.Title)
	if err != nil {
		c.Mess <- &modals.Update{Error: fmt.Sprintf("error while creating chat: %v", err)}
		return
	}

	fmt.Println("chat created: ", chat.Title)

	//c.Mess <- &modals.Update{NewChatCreated: chat}

	DVSr.mu.Lock()
	newChannel := NewChannel(chat.ID)
	DVSr.Channels[chat.ID] = newChannel
	newChannel.Users[c] = true
	go newChannel.Run()
	DVSr.mu.Unlock()
	// Start the channel to handle the chat
	HandleAllJoin(chat)
}

// HandleAllJoin adds all users to the channel
func HandleAllJoin(chat *modals.Chat) {
	if channel, ok := DVSr.Channels[chat.ID]; ok {
		for _, parti := range chat.Members {
			fmt.Println("user", parti.UserID)
			if client, ok := DVSr.Users[parti.UserID]; ok {
				client.Join <- chat.ID
				channel.Join <- client
				client.Mess <- &modals.Update{NewChatCreated: chat}
			} else {
				fmt.Println("user not active", parti.UserID)
			}
		}
	} else {
		log.Fatalln("channel not found")
	}
}