package server

import (
	"encoding/json"
	"fmt"
	"gRMS/modals"
)

type Req struct {
	Message     *MessQuery    `json:"message"`
	NewChat     *NewChatQuery `json:"new_chat"`
	ChatJoin    *UserQuery    `json:"add_user"`
	ChatKick    *UserQuery    `json:"kick_user"`
	Promote     *UserQuery    `json:"promote"`
	Demote      *UserQuery    `json:"demote"`
	ChangeTitle *ChatQuery    `json:"change_title"`
	GetUser     uint64        `json:"get_user"`
	GetChat     uint64        `json:"get_chat"`
	GetSelf     uint64        `json:"get_self"`
	LeaveChat   uint64        `json:"leave_chat"`
	//Forward *msgService.ForwardQuery  `json:"forward"`
}

func (sr *dvs) HandleReq(c Client, p []byte) {
	//fmt.Println("handling request", string(p))
	req := &Req{}
	err := json.Unmarshal(p, req)
	if err != nil {
		c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error unmarshaling request: %v", err))
		return
	}

	switch {
	case req.Message != nil:
		//fmt.Println("message received", req.Message.Text)
		sr.HandleMess(c, req.Message)
	case req.NewChat != nil:
		sr.HandleNewChat(req.NewChat, c)
	case req.ChatJoin != nil:
		sr.HandleAddToChat(c, req.ChatJoin)
	case req.GetUser != 0:
		user, err := sr.Dbs.GetUser(req.GetUser)
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding user: %v", err))
		} else {
			c.Updates() <- modals.UserUpdate(user)
		}
	case req.GetChat != 0:
		chat, err := sr.Dbs.GetChat(req.GetChat)
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding chat: %v", err))
		} else {
			c.Updates() <- modals.ChatUpdate(chat)
		}
	case req.GetSelf != 0:
		user, err := sr.Dbs.GetUser(c.GetUserID())
		if err != nil {
			c.Updates() <- modals.ErrorUpdate(fmt.Sprintf("error finding user: %v", err))
		} else {
			c.Updates() <- &modals.Update{Self: user}
		}
	}
}