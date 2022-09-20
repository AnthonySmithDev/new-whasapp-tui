package repository

import (
	"strings"

	db "github.com/sonyarouje/simdb"
	"go.mau.fi/whatsmeow/binary/proto"
)

type Conversation struct {
	*proto.Conversation
}

func (c Conversation) ID() (jsonField string, value interface{}) {
	value = c.GetId()
	jsonField = "conversation_id"
	return
}

func (c Conversation) Title() string       { return c.GetName() }
func (c Conversation) Description() string { return c.GetId() }
func (c Conversation) FilterValue() string { return c.GetId() }

func (c Conversation) IsGroup() bool {
	return strings.Contains(c.GetId(), "@g.us")
}

type ConversationInter interface {
	Create(Conversation)
	FindOne() Conversation
	FindMany() []Conversation
}

type conversationImpl struct {
	db *db.Driver
}

func (repository *conversationImpl) Create(conversation Conversation) {
	repository.db.Open(Conversation{}).Insert(conversation)
}

func (repository *conversationImpl) FindOne() Conversation {
	var conversation Conversation
	repository.db.Open(Conversation{}).First().AsEntity(&conversation)
	return conversation
}

func (repository *conversationImpl) FindMany() []Conversation {
	var conversation []Conversation
	repository.db.Open(Conversation{}).Get().AsEntity(&conversation)
	return conversation
}

func NewConversation(db *db.Driver) ConversationInter {
	return &conversationImpl{db}
}
