package modals

type File struct {
	ID       uint64   `json:"id"`
	Filesize uint64   `json:"filesize"`
	Filepath string   `json:"filepath"`
	Filetype Filetype `json:"filetype"`
}

type Filetype byte

const (
	PhotoType Filetype = iota
	VideoType
	AudioType
	DocumentType
	StickerType
	AnimationType
)