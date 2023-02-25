package msgService

// VideoQuery is query format for sending video
type VideoQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the video to be sent
	VideoID uint64 `json:"video"`
	// Caption is the video caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// TextQuery is query format for sending message
type TextQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// Text the body of the text message
	Text string `json:"text"`
	// ReplyToMessageId is the id of the replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// StickerQuery is query format for sending video
type StickerQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the video to be sent
	StickerID uint64 `json:"video"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// PhotoQuery is query format for sending photo
type PhotoQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// PhotoID is the file ID of the photo to be sent
	PhotoID uint64 `json:"photo"`
	// Caption is the photo caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// DocumentQuery is the query format for sending document
type DocumentQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// DocumentID is the file ID of the photo to be sent
	DocumentID uint64 `json:"document"`
	// Caption is the document caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// AudioQuery is query format for sending audio
type AudioQuery struct {
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AudioID is the file ID of the photo to be sent
	AudioID uint64 `json:"audio"`
	// Caption is the audio caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}

// AnimationQuery is query format for sending animation
type AnimationQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AnimationID is the file ID of the photo to be sent
	AnimationID uint64 `json:"animation"`
	// Caption is the animation caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail of the animation
	Thumb uint64 `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}