package dbService

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
		file.Filepath = ph.GetMetaData().Filepath
	case modals.AudioType:
		au, err := sr.GetAudio(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = au.GetMetaData().Filepath
	case modals.VideoType:
		vd, err := sr.GetVideo(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = vd.GetMetaData().Filepath
	case modals.DocumentType:
		doc, err := sr.GetDocument(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = doc.GetMetaData().Filepath
	case modals.StickerType:
		stk, err := sr.GetSticker(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = stk.GetMetaData().Filepath
	case modals.AnimationType:
		anim, err := sr.GetAnimation(fileId)
		if err != nil {
			return nil, err
		}
		file.Filepath = anim.GetMetaData().Filepath
	}

	return &file, nil
}