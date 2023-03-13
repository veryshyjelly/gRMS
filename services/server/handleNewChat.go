package server

import (
	"fmt"
	"gRMS/modals"
	"log"
)

type NewChatQuery struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

// HandleNewChat creates a new chat and starts a new channel to handle it
func (sr *dvs) HandleNewChat(chatQuery *NewChatQuery, c Client) {
	users := []uint64{c.GetUserID()}
	// Check if the participants are valid
	if chatQuery.Participants == nil {
		c.Updates() <- modals.ErrorUpdate("error: participants list not present")
		return
	}
	// Loop through the participants find their userID and add them to users list
	for _, username := range chatQuery.Participants {
		user, err := sr.findUser(username)
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding user: %v", err))
			return
		}
		users = append(users, user.ID)
	}

	chat, err := sr.createChat(users, chatQuery.Title)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error creating chat: %v", err))
		return
	}

	newChannel := NewChannel(chat.ID, c)
	go newChannel.Run(sr)
	// Start the channel to handle the chat

	sr.LockChannels()
	sr.AddChannel() <- newChannel

	sr.HandleAllJoin(chat)
}

// HandleAllJoin adds all users to the channel
func (sr *dvs) HandleAllJoin(chat *modals.Chat) {
	sr.LockChannels()
	defer sr.UnlockChannels()

	if channel, ok := sr.ActiveChannels()[chat.ID]; ok {
		for _, parti := range chat.Members {
			if client, ok := sr.ActiveUsers()[parti.UserID]; ok {
				client.ChatJoin() <- chat.ID
				channel.UserJoin() <- client
				client.Updates() <- modals.NewChatUpdate(chat)
			} else {
				fmt.Println("user not active", parti.UserID)
			}
		}
	} else {
		log.Fatalln("channel not found")
	}
}