package dbservice

import "chat-app/modals"

// NewFile function should be used by the getFile method to get the downloadable file path
func (sr *DBService) NewFile(fileId uint64, filetype modals.Filetype) (*modals.File, error) {
	file := modals.File{ID: fileId, Filetype: filetype}

	switch filetype {
	case modals.PhotoType:
		ph, err := sr.GetPhoto(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = ph.GetFilepath()
	case modals.AudioType:
		au, err := sr.GetAudio(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = au.GetFilepath()
	case modals.VideoType:
		vd, err := sr.GetVideo(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = vd.GetFilepath()
	case modals.DocumentType:
		doc, err := sr.GetDocument(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = doc.GetFilepath()
	case modals.StickerType:
		stk, err := sr.GetSticker(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = stk.GetFilepath()
	case modals.AnimationType:
		anim, err := sr.GetAnimation(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = anim.GetFilepath()
	}

	return &file, nil
}