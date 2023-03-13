package server

import (
	"fmt"
	"gRMS/modals"
	dbservice "gRMS/services/db"
	msgService "gRMS/services/msg"
	"github.com/gofiber/websocket/v2"
	"sync"
)

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
	NewClient(*modals.User, *websocket.Conn) Client
}

type dvs struct {
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

func NewDvService(mgs msgService.MsgS, dbs dbservice.DBS) DVS {
	return &dvs{
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

func (sr *dvs) Run() {
	fmt.Println("starting dv service")
	defer fmt.Println("stopping dv service")
	for {
		select {
		case ch := <-sr.NewChannel:
			fmt.Println("new channel active", ch)
			sr.Channels[ch.GetChatID()] = ch
			sr.UnlockChannels()
		case cl := <-sr.NewUser:
			fmt.Println("new user active", cl.GetUsername())
			sr.Users[cl.GetUserID()] = cl
			sr.UnlockUsers()
		case chatID := <-sr.EndChannel:
			fmt.Println("stopping chanel", chatID)
			delete(sr.Channels, chatID)
			sr.UnlockChannels()
		case userID := <-sr.UserLeft:
			fmt.Println("user left", userID)
			delete(sr.Users, userID)
			sr.UnlockUsers()
		}
	}
}

func (sr *dvs) AddChannel() chan Channel {
	return sr.NewChannel
}

func (sr *dvs) ActiveChannels() map[uint64]Channel {
	return sr.Channels
}

func (sr *dvs) StopChannel() chan uint64 {
	return sr.EndChannel
}

func (sr *dvs) AddUser() chan Client {
	return sr.NewUser
}

func (sr *dvs) ActiveUsers() map[uint64]Client {
	return sr.Users
}

func (sr *dvs) LeaveUser() chan uint64 {
	return sr.UserLeft
}

func (sr *dvs) LockUsers() {
	sr.muUser.Lock()
}

func (sr *dvs) UnlockUsers() {
	sr.muUser.Unlock()
}

func (sr *dvs) LockChannels() {
	sr.muChannel.Lock()
}

func (sr *dvs) UnlockChannels() {
	sr.muChannel.Unlock()
}