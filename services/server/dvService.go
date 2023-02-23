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
	Run()
	Lock()
	Unlock()
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

func (dvs *DvService) Lock() {
	fmt.Println("locking dv service")
	dvs.mu.Lock()
}

func (dvs *DvService) Unlock() {
	fmt.Println("unlocking dv service")
	dvs.mu.Unlock()
}

func NewDvService(mgs msgService.MsgS) *DvService {
	return &DvService{
		Mgs:        mgs,
		Channels:   make(map[uint64]*Channel),
		Users:      make(map[uint64]*Client),
		NewChannel: make(chan *Channel),
		NewUser:    make(chan *Client),
		EndChannel: make(chan uint64),
		UserLeft:   make(chan uint64),
	}
}

func (dvs *DvService) Run() {
	fmt.Println("starting dv service")
	for {
		select {
		case channel := <-dvs.NewChannel:
			fmt.Println("new channel active", channel.ChatID)
			dvs.Channels[channel.ChatID] = channel
			DVSr.Unlock()
		case client := <-dvs.NewUser:
			fmt.Println("new user active", client.GetUsername())
			dvs.Users[client.GetUserID()] = client
			client.ID = uint64(len(dvs.Users))
			DVSr.Unlock()
		case chatID := <-dvs.EndChannel:
			fmt.Println("stopping chanel", chatID)
			delete(dvs.Channels, chatID)
		case userID := <-dvs.UserLeft:
			fmt.Println("user left", userID)
			delete(dvs.Users, userID)
		}
	}
	fmt.Println("stopping dv service")
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