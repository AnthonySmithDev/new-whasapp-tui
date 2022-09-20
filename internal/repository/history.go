package repository

import (
	"strings"
	"time"

	"go.mau.fi/whatsmeow/binary/proto"
)

func CreateHistory(convRepo ConversationInter, msgsRepo MessageInter, proto []*proto.Conversation) {
	for _, conversation := range proto {
		new := Conversation{
			JID:     conversation.GetId(),
			Name:    conversation.GetName(),
			IsGroup: isGroup(conversation.GetId()),
		}
		convRepo.Create(new)
		messages := convertMessage(conversation.GetMessages())
		msgsRepo.CreateBatch(messages)
	}
}

func convertMessage(proto []*proto.HistorySyncMsg) []Message {
	var messages []Message
	for _, message := range proto {
		new := Message{
			MsgOrderId:       message.GetMsgOrderId(),
			Msg:              getMessageText(message.GetMessage().GetMessage()),
			MsgType:          getTypeMessage(message.GetMessage().GetMessage()),
			FromMe:           message.GetMessage().GetKey().GetFromMe(),
			RemoteJid:        message.GetMessage().GetKey().GetRemoteJid(),
			MessageTimestamp: time.Unix(int64(message.GetMessage().GetMessageTimestamp()), 0),
		}
		messages = append(messages, new)
	}
	return messages
}

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