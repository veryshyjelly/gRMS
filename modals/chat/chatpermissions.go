package chatmodals

type ChatPermissions struct {
	CanSendMessages  *bool `json:"can_send_messages,omitempty"`
	CanSendAudios    *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos    *bool `json:"can_send_photos,omitempty"`
	CanSendVideos    *bool `json:"can_send_videos,omitempty"`
	CanInviteUsers   *bool `json:"can_invite_users,omitempty"`
	CanPinMessages   *bool `json:"can_pin_messages,omitempty"`
}