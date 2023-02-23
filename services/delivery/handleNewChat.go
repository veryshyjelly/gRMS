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
	users := []uint64{c.user.ID}
	// Check if the participants are valid
	if chatQuery.Participants == nil {
		c.Updates() <- modals.ErrorUpdate("error: participants list not present")
		return
	}
	// Loop through the participants find their userID and add them to users list
	for _, username := range chatQuery.Participants {
		user, err := dbService.DBSr.FindUser(username)
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding user: %v", err))
			return
		}
		users = append(users, user.ID)
	}

	chat, err := dbService.DBSr.CreateChat(users, chatQuery.Title)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error creating chat: %v", err))
		return
	}

	fmt.Println("chat created chat: ", chat.Title)

	newChannel := NewChannel(chat.ID, c)
	go newChannel.Run()
	// Start the channel to handle the chat

	DVSr.AddChannel() <- newChannel
	HandleAllJoin(chat)
}

// HandleAllJoin adds all users to the channel
func HandleAllJoin(chat *modals.Chat) {
	if channel, ok := DVSr.ActiveChannels()[chat.ID]; ok {
		for _, parti := range chat.Members {
			fmt.Println("user", parti.UserID)
			if client, ok := DVSr.ActiveUsers()[parti.UserID]; ok {
				client.ChatJoin() <- chat.ID
				channel.Join <- client
				client.Updates() <- &modals.Update{NewChatCreated: chat}
			} else {
				fmt.Println("user not active", parti.UserID)
			}
		}
	} else {
		log.Fatalln("channel not found")
	}
}