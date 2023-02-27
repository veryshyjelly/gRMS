package server

import (
	"chat-app/modals"
	dbservice "chat-app/services/db"
	msgService "chat-app/services/msg"
	"fmt"
	"sync"
)

var DVSr DVS

type DVS interface {
	SendMess(*modals.Message)
	AddChannel() chan Channel
	ActiveChannels() map[uint64]Channel
	AddUser() chan Client
	ActiveUsers() map[uint64]Client
	StopChannel() chan uint64
	LeaveUser() chan uint64
	HandleMess(Client, *MessQuery)
	HandleReq(Client, []byte)
	Run()
	LockUsers()
	UnlockUsers()
	LockChannels()
	UnlockChannels()
}

type DvService struct {
	Mgs        msgService.MsgS
	Dbs        dbservice.DBS
	Channels   map[uint64]Channel
	Users      map[uint64]Client
	NewChannel chan Channel
	NewUser    chan Client
	EndChannel chan uint64
	UserLeft   chan uint64
	muUser     sync.Mutex
	muChannel  sync.Mutex
}

func NewDvService(mgs msgService.MsgS, dbs dbservice.DBS) *DvService {
	return &DvService{
		Mgs:        mgs,
		Dbs:        dbs,
		Channels:   make(map[uint64]Channel),
		Users:      make(map[uint64]Client),
		NewChannel: make(chan Channel),
		NewUser:    make(chan Client),
		EndChannel: make(chan uint64),
		UserLeft:   make(chan uint64),
	}
}

func (dvs *DvService) Run() {
	fmt.Println("starting dv service")
	defer fmt.Println("stopping dv service")
	for {
		select {
		case ch := <-dvs.NewChannel:
			fmt.Println("new channel active", ch)
			dvs.Channels[ch.GetChatID()] = ch
			dvs.UnlockChannels()
		case cl := <-dvs.NewUser:
			fmt.Println("new user active", cl.GetUsername())
			dvs.Users[cl.GetUserID()] = cl
			dvs.UnlockUsers()
		case chatID := <-dvs.EndChannel:
			fmt.Println("stopping chanel", chatID)
			delete(dvs.Channels, chatID)
			dvs.UnlockChannels()
		case userID := <-dvs.UserLeft:
			fmt.Println("user left", userID)
			delete(dvs.Users, userID)
			dvs.UnlockUsers()
		}
	}
}

func (dvs *DvService) AddChannel() chan Channel {
	return dvs.NewChannel
}

func (dvs *DvService) ActiveChannels() map[uint64]Channel {
	return dvs.Channels
}

func (dvs *DvService) StopChannel() chan uint64 {
	return dvs.EndChannel
}

func (dvs *DvService) AddUser() chan Client {
	return dvs.NewUser
}

func (dvs *DvService) ActiveUsers() map[uint64]Client {
	return dvs.Users
}

func (dvs *DvService) LeaveUser() chan uint64 {
	return dvs.UserLeft
}

func (dvs *DvService) LockUsers() {
	fmt.Println("locking users")
	dvs.muUser.Lock()
}

func (dvs *DvService) UnlockUsers() {
	fmt.Println("unlocking users")
	dvs.muUser.Unlock()
}

func (dvs *DvService) LockChannels() {
	fmt.Println("locking channel")
	dvs.muChannel.Lock()
}

func (dvs *DvService) UnlockChannels() {
	fmt.Println("unlocking channel")
	dvs.muChannel.Unlock()
}