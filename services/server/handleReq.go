package server

import (
	"chat-app/modals"
	dbService "chat-app/services/db"
	"encoding/json"
	"fmt"
)

type Req struct {
	Message  *MessQuery    `json:"message"`
	NewChat  *NewChatQuery `json:"new_chat"`
	ChatJoin *AddUserQuery `json:"add_user"`
	GetUser  uint64        `json:"get_user"`
	GetChat  uint64        `json:"get_chat"`
	//Forward *msgService.ForwardQuery  `json:"forward"`
}

func (dvs *DvService) HandleReq(c Client, p []byte) {
	fmt.Println("handling request", string(p))
	req := &Req{}
	err := json.Unmarshal(p, req)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error unmarshaling request: %v", err))
		return
	}

	switch {
	case req.Message != nil:
		fmt.Println("message received", req.Message.Text)
		DVSr.HandleMess(c, req.Message)
	case req.NewChat != nil:
		HandleNewChat(req.NewChat, c)
	case req.ChatJoin != nil:
		HandleAddToChat(c, req.ChatJoin)
	case req.GetUser != 0:
		user, err := dbService.DBSr.GetUser(req.GetUser)
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding user: %v", err))
		} else {
			c.Updates() <- modals.UserUpdate(user)
		}
	}
}