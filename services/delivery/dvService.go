package delivery

import (
	"chat-app/modals"
	msgService "chat-app/services/msg"
	"log"
	"sync"
)

var DVSr *DvService

type DVS interface {
	Send(*modals.Message)
}

type DvService struct {
	Mgs      msgService.MsgS
	Channels map[uint64]*Channel
	mu       sync.Mutex
}

func NewDvService(mgs msgService.MsgS) *DvService {
	return &DvService{
		Mgs: mgs,
	}
}

func (dvs *DvService) Send(msg *modals.Message) {
	if channel, ok := dvs.Channels[msg.Chat.ID]; ok {
		channel.Mess <- msg
	} else {
		log.Println("error channel not found")
	}
}