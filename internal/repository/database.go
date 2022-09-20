package repository

import (
	db "github.com/sonyarouje/simdb"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type DB struct {
	Message      MessageInter
	Conversation ConversationInter
}

func NewDB() *DB {
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	return &DB{
		Message:      NewMessage(driver),
		Conversation: NewConversation(driver),
	}
}

func (db *DB) CreateHistory(history *proto.HistorySync, cli *whatsmeow.Client) {
	for _, conv := range history.GetConversations() {
		chatJID, _ := types.ParseJID(conv.GetId())
		for _, historyMsg := range conv.GetMessages() {
			msg, _ := cli.ParseWebMessage(chatJID, historyMsg.GetMessage())
			db.Message.Create(Message{msg})
		}
		conv.Messages = nil
		db.Conversation.Create(Conversation{conv})
	}
}

func (db *DB) CreateMessage(msg *events.Message) {
	db.Message.Create(Message{msg})
}
