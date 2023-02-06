package modals

type File struct {
	ID       uint64 `json:"id"`
	Filesize uint64 `json:"filesize"`
	Filepath string `json:"filepath"`
}

func NewFile() *File {
	return &File{}
}