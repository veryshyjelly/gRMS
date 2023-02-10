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

// NewFile function should be used by the getFile method to get the downloadable file path
func (sr *DBService) NewFile(fileId uint64, filetype Filetype) (*File, error) {
	file := File{ID: fileId, Filetype: filetype}

	switch filetype {
	case PhotoType:
		ph, err := sr.GetPhoto(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = ph.GetFilepath()
	case AudioType:
		au, err := sr.GetAudio(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = au.GetFilepath()
	case VideoType:
		vd, err := sr.GetVideo(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = vd.GetFilepath()
	case DocumentType:
		doc, err := sr.GetDocument(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = doc.GetFilepath()
	case StickerType:
		stk, err := sr.GetSticker(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = stk.GetFilepath()
	case AnimationType:
		anim, err := sr.GetAnimation(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = anim.GetFilepath()
	}

	return &file, nil
}