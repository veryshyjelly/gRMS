package msgService

import (
	"chat-app/modals"
	dbService "chat-app/services/db"
)

var MGSr *MsgService

type MsgS interface {
	Text(*TextQuery) (*modals.Message, error)
	Photo(*PhotoQuery) (*modals.Message, error)
	Video(*VideoQuery) (*modals.Message, error)
	Audio(*AudioQuery) (*modals.Message, error)
	Document(*DocumentQuery) (*modals.Message, error)
	Sticker(*StickerQuery) (*modals.Message, error)
	Animation(*AnimationQuery) (*modals.Message, error)
}

type MsgService struct {
	dbs dbService.DBS
}

func NewMsgService(db dbService.DBS) *MsgService {
	return &MsgService{dbs: db}
}