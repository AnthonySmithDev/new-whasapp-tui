package repository

import (
	"gorm.io/gorm"
)

type Conversation struct {
	JID     string
	Name    string
	IsGroup bool
}

type ConversationInter interface {
	Create(Conversation)
	CreateBatch([]Conversation)
	FindOne() Conversation
	FindAll() []Conversation
}

type conversationImpl struct {
	db *gorm.DB
}

func (repository *conversationImpl) Create(conversation Conversation) {
	repository.db.Create(&conversation)
}

func (repository *conversationImpl) CreateBatch(conversations []Conversation) {
	repository.db.Create(&conversations)
}

func (repository *conversationImpl) FindOne() Conversation {
	var conversation Conversation
	repository.db.First(&conversation)
	return conversation
}

func (repository *conversationImpl) FindAll() []Conversation {
	var conversation []Conversation
	repository.db.Find(&conversation)
	return conversation
}

func NewConversation(db *gorm.DB) ConversationInter {
	db.AutoMigrate(&Conversation{})
	return &conversationImpl{db}
}
func (c *Conversation) AfterFind(tx *gorm.DB) (err error) {
	// if !c.IsGroup {
	// 	jid, _ := types.ParseJID(c.JID)
	// 	contact, _ := wa.Store.GetContact(jid)
	// 	c.Name = contact.FullName
	// }
	return
}
