package modals

import (
	"gorm.io/gorm"
)

type File struct {
	ID       uint64 `json:"id"`
	Filesize uint64 `json:"filesize"`
	Filepath string `json:"filepath"`
	Filetype Filetype
}

type Filetype byte

const (
	PhotoType    Filetype = 1
	AudioType             = 2
	VideoType             = 3
	DocumentType          = 4
	StickerType           = 5
)

// NewFile function should be used by the getFile method to get the downloadable file path
func NewFile(db *gorm.DB, fileId uint64, filetype Filetype) (*File, error) {
	file := File{ID: fileId, Filetype: filetype}
	switch filetype {
	case PhotoType:
		ph, err := FindPhoto(db, fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = ph.Filepath
	case AudioType:
		au, err := FindAudio(db, fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = au.Filepath
	case VideoType:
		vd, err := FindVideo(db, fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = vd.Filepath
	case DocumentType:
		doc, err := FindDocument(db, fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = doc.Filepath
	case StickerType:
		stk, err := FindSticker(db, fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = stk.Filepath
	}

	return &file, nil
}