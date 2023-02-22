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
	Leave    chan uint64
	Stop     chan uint64
}

func NewDvService(mgs msgService.MsgS) *DvService {
	return &DvService{
		Mgs:      mgs,
		Channels: make(map[uint64]*Channel),
		Users:    make(map[uint64]*Client),
		Leave:    make(chan uint64),
		Stop:     make(chan uint64),
	}
}