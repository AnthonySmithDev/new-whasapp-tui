package repository

import (
	"strings"

	"go.mau.fi/whatsmeow/binary/proto"
)

func isGroup(jid string) bool {
	return strings.Contains(jid, "@g.us")
}

func getTypeMessage(message *proto.Message) MsgType {
	if message.GetExtendedTextMessage() != nil {
		return Link
	}
	if message.GetImageMessage() != nil {
		return Image
	}
	if message.GetAudioMessage() != nil {
		return Audio
	}
	if message.GetVideoMessage() != nil {
		return Video
	}
	if message.GetStickerMessage() != nil {
		return Sticker
	}
	return Text
}

func getMessageText(message *proto.Message) string {
	if message.GetExtendedTextMessage() != nil {
		return "Link"
	}
	if message.GetImageMessage() != nil {
		return "Image"
	}
	if message.GetAudioMessage() != nil {
		return "Audio"
	}
	if message.GetVideoMessage() != nil {
		return "Video"
	}
	if message.GetStickerMessage() != nil {
		return "Sticker"
	}
	return message.GetConversation()
}