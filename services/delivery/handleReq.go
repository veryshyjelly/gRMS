package delivery

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

func (c *Client) HandleReq(p []byte) {
	req := &Req{}
	err := json.Unmarshal(p, req)
	if err != nil {
		e := fmt.Sprintf("error while unmarshaling message: %v", err)
		c.Mess <- modals.NewUpdate(0, &modals.Message{Text: &e})
		return
	}

	if req.Message != nil {
		fmt.Println("message received", req.Message.Text)
		c.HandleMess(req.Message)
	}

	if req.NewChat != nil {
		c.HandleNewChat(req.NewChat)
	}

	if req.ChatJoin != nil {
		c.HandleAddToChat(req.ChatJoin)
	}

	if req.GetUser != 0 {
		user, err := dbService.DBSr.GetUser(req.GetUser)
		if err != nil {
			c.Mess <- &modals.Update{Error: fmt.Sprintf("error finding user: %v", err)}
		} else {
			c.Mess <- &modals.Update{User: user}
		}
	}
}