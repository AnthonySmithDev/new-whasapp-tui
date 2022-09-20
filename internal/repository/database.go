package repository

import (
	db "github.com/sonyarouje/simdb"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type DB struct {
	message      MessageInter
	conversation ConversationInter
}

func NewDB() *DB {
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	return &DB{
		message:      NewMessage(driver),
		conversation: NewConversation(driver),
	}
}

func (db *DB) CreateHistory(history *proto.HistorySync, cli *whatsmeow.Client) {
	for _, conv := range history.GetConversations() {
		chatJID, _ := types.ParseJID(conv.GetId())
		db.conversation.Create(Conversation{conv})
		for _, historyMsg := range conv.GetMessages() {
			msg, _ := cli.ParseWebMessage(chatJID, historyMsg.GetMessage())
			db.message.Create(Message{msg.Message})
		}
	}
}
