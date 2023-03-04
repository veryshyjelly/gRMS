package server

import (
	"fmt"
	"gRMS/modals"
	dbService "gRMS/services/db"
	"log"
)

type NewChatQuery struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

// HandleNewChat creates a new chat and starts a new channel to handle it
func HandleNewChat(chatQuery *NewChatQuery, c Client) {
	users := []uint64{c.GetUserID()}
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

	newChannel := NewChannel(chat.ID, c)
	go newChannel.Run()
	// Start the channel to handle the chat

	DVSr.LockChannels()
	DVSr.AddChannel() <- newChannel

	HandleAllJoin(chat)
}

// HandleAllJoin adds all users to the channel
func HandleAllJoin(chat *modals.Chat) {
	DVSr.LockChannels()
	defer DVSr.UnlockChannels()

	if channel, ok := DVSr.ActiveChannels()[chat.ID]; ok {
		for _, parti := range chat.Members {
			if client, ok := DVSr.ActiveUsers()[parti.UserID]; ok {
				client.ChatJoin() <- chat.ID
				channel.UserJoin() <- client
				client.Updates() <- &modals.Update{NewChatCreated: chat}
			} else {
				fmt.Println("user not active", parti.UserID)
			}
		}
	} else {
		log.Fatalln("channel not found")
	}
}
