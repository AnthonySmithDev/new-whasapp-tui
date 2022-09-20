package repository

import (
	db "github.com/sonyarouje/simdb"
	"go.mau.fi/whatsmeow/binary/proto"
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
	*proto.Message
}

func (c Message) ID() (jsonField string, value interface{}) {
	value = c.GetOrderMessage().GetOrderId()
	jsonField = "message_id"
	return
}

type MessageInter interface {
	Create(Message)
	FindOne(string) Message
	FindMany(string) []Message
}

type MessageImpl struct {
	db *db.Driver
}

func (repository *MessageImpl) Create(message Message) {
	repository.db.Open(Message{}).Insert(message)
}

func (repository *MessageImpl) FindOne(jid string) Message {
	var message Message
	repository.db.Open(Message{}).Where("conversation_id", "=", jid).First().AsEntity(&message)
	return message
}

func (repository *MessageImpl) FindMany(jid string) []Message {
	var messages []Message
	repository.db.Open(Message{}).Where("conversation_id", "=", jid).Get().AsEntity(&messages)
	return messages
}

func NewMessage(db *db.Driver) MessageInter {
	return &MessageImpl{db}
}
