package modals

type File struct {
	ID       uint64 `json:"id"`
	Filesize uint64 `json:"filesize"`
	Filepath string `json:"filepath"`
	Filetype Filetype
}

type Filetype byte

const (
	PhotoType     Filetype = 1
	VideoType              = 2
	AudioType              = 3
	DocumentType           = 4
	StickerType            = 5
	AnimationType          = 6
)