package modals

type MediaMD struct {
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filepath is the path of the file
	Filepath string `json:"filepath"`
}