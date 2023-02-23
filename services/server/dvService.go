package server

import (
	"chat-app/modals"
	msgService "chat-app/services/msg"
	"fmt"
	"sync"
)

var DVSr DVS

type DVS interface {
	SendMess(*modals.Message)
	SendChat(*modals.Chat)
	AddChannel() chan *Channel
	ActiveChannels() map[uint64]*Channel
	AddUser() chan *Client
	ActiveUsers() map[uint64]*Client
	StopChannel() chan uint64
	LeaveUser() chan uint64
	HandleMess(*MessQuery, *Client)
	HandleReq([]byte, *Client)
}

type DvService struct {
	Mgs        msgService.MsgS
	Channels   map[uint64]*Channel
	Users      map[uint64]*Client
	NewChannel chan *Channel
	NewUser    chan *Client
	EndChannel chan uint64
	UserLeft   chan uint64
	mu         sync.Mutex
}

func NewDvService(mgs msgService.MsgS) *DvService {
	return &DvService{
		Mgs:        mgs,
		Channels:   make(map[uint64]*Channel),
		Users:      make(map[uint64]*Client),
		NewChannel: make(chan *Channel, 100),
		NewUser:    make(chan *Client, 100),
		EndChannel: make(chan uint64, 100),
		UserLeft:   make(chan uint64, 100),
	}
}

func (dvs *DvService) Run() {
	for {
		select {
		case channel := <-dvs.NewChannel:
			dvs.Channels[channel.ChatID] = channel
		case user := <-dvs.NewUser:
			fmt.Println("new user active", user.GetUsername())
			dvs.Users[user.ID] = user
			user.ID = uint64(len(dvs.Users))
		case chatID := <-dvs.EndChannel:
			fmt.Println("stopping chanel", chatID)
			delete(dvs.Channels, chatID)
		case userID := <-dvs.UserLeft:
			fmt.Println("user left", userID)
			delete(dvs.Users, userID)
		}
	}
}

func (dvs *DvService) AddChannel() chan *Channel {
	return dvs.NewChannel
}

func (dvs *DvService) ActiveChannels() map[uint64]*Channel {
	return dvs.Channels
}

func (dvs *DvService) StopChannel() chan uint64 {
	return dvs.EndChannel
}

func (dvs *DvService) AddUser() chan *Client {
	return dvs.NewUser
}

func (dvs *DvService) ActiveUsers() map[uint64]*Client {
	return dvs.Users
}

func (dvs *DvService) LeaveUser() chan uint64 {
	return dvs.UserLeft
}