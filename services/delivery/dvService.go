package delivery

import (
	"chat-app/modals"
	msgService "chat-app/services/msg"
	"sync"
)

var DVSr *DvService

type DVS interface {
	SendMess(*modals.Message)
	SendChat(*modals.Chat)
}

type DvService struct {
	Mgs      msgService.MsgS
	Channels map[uint64]*Channel
	Users    map[uint64]*Client
	mu       sync.Mutex
}

func NewDvService(mgs msgService.MsgS) *DvService {
	return &DvService{
		Mgs:      mgs,
		Channels: make(map[uint64]*Channel),
		Users:    make(map[uint64]*Client),
	}
}