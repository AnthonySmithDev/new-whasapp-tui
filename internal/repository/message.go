package repository

import (
	"github.com/charmbracelet/lipgloss"
	db "github.com/sonyarouje/simdb"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

type MsgType uint

const (
	Text MsgType = iota
	Link
	Audio
	Image
	Video
	Sticker
)

type Message struct {
	*events.Message
}

func (c Message) ID() (jsonField string, value interface{}) {
	value = c.Info.ID
	jsonField = "message_id"
	return
}

func (message *Message) ToString() string {
	return messageToText(message.Message.Message)
}

type MessageInter interface {
	Create(Message)
	FindOne(string) Message
	FindMany(string) Messages
}

type MessageImpl struct {
	db *db.Driver
}

type Messages []Message

func (messages Messages) ToList() []string {
	var list []string
	for _, message := range messages {
		list = append(list, messageToText(message.Message.Message))
	}
	return list
}

func (repository *MessageImpl) Create(message Message) {
	repository.db.Open(Message{}).Insert(message)
}

func (repository *MessageImpl) FindOne(jid string) Message {
	var messages []Message
	repository.db.Open(Message{}).Where("Info.Chat", "=", jid).Get().AsEntity(&messages)
	return messages[len(messages)-1]
}

func (repository *MessageImpl) FindMany(jid string) Messages {
	var messages []Message
	repository.db.Open(Message{}).Where("Info.Chat", "=", jid).Get().AsEntity(&messages)
	return messages
}

func reverse(messages []Message) []Message {
	newMessages := make([]Message, len(messages))
	for i, j := 0, len(messages)-1; i <= j; i, j = i+1, j-1 {
		newMessages[i], newMessages[j] = messages[j], messages[i]
	}
	return newMessages
}

func NewMessage(db *db.Driver) MessageInter {
	return &MessageImpl{db}
}

func messageToText(message *proto.Message) string {
	if message.GetExtendedTextMessage() != nil {
		return "Link"
	}
	if message.GetImageMessage() != nil {
		return "Image"
	}
	if message.GetDocumentMessage() != nil {
		return "Document"
	}
	if message.GetContactMessage() != nil {
		return "Contact"
	}
	if message.GetLocationMessage() != nil {
		return "Location"
	}
	if message.GetInvoiceMessage() != nil {
		return "Invoice"
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

func messageToStyle(message *proto.Message) lipgloss.Style {
	if message.GetExtendedTextMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	}
	if message.GetImageMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	}
	if message.GetDocumentMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	}
	if message.GetContactMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	}
	if message.GetLocationMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	}
	if message.GetInvoiceMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	}
	if message.GetAudioMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	}
	if message.GetVideoMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	}
	if message.GetStickerMessage() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
}
