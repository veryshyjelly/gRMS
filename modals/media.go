package modals

import "time"

type MediaMD struct {
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filepath is the path of the file
	Filepath  string `json:"filepath"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}