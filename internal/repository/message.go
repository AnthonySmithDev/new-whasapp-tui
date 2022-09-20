package repository

import (
	"time"

	"gorm.io/gorm"
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
	MessageTimestamp time.Time
	MsgOrderId       uint64
	RemoteJid        string
	FromMe           bool
	Msg              string
	MsgType          MsgType
}

type MessageInter interface {
	Create(Message)
	CreateBatch([]Message)
	FindOne(string) Message
	FindAll(string) []Message
}

type messageImpl struct {
	db *gorm.DB
}

func (repository *messageImpl) Create(message Message) {
	repository.db.Create(&message)
}

func (repository *messageImpl) CreateBatch(messages []Message) {
	repository.db.Create(&messages)
}

func (repository *messageImpl) FindAll(jid string) []Message {
	var message []Message
	repository.db.Where(&Message{RemoteJid: jid}).Find(&message)
	return message
}

func (repository *messageImpl) FindOne(jid string) Message {
	var message Message
	repository.db.Where(&Message{RemoteJid: jid}).Last(&message)
	return message
}

func NewMessage(db *gorm.DB) MessageInter {
	db.AutoMigrate(&Message{})
	return &messageImpl{db}
}
