package delivery

import (
	"chat-app/modals"
	"encoding/json"
	"fmt"
)

type Req struct {
	Message  *MessQuery     `json:"message"`
	NewChat  *NewChatQuery  `json:"new_chat"`
	ChatJoin *ChatJoinQuery `json:"chat_join"`
	GetUser  uint64         `json:"get_user"`
	GetChat  uint64         `json:"get_chat"`
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
}