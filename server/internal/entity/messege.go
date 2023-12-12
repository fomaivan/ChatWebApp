package entity

import "time"

type Message struct {
	MessageId uint `gorm:"primaryKey"`
	InboxId   uint
	UserId    uint
	Content   string
	SendTime  time.Time
}

type SendMessage struct {
	From    uint
	To      uint
	Content string
}

type MessageRepository interface {
	Create(*Message) error
	Update(*Message) error
	Delete(id uint) error
	GetAll() (*[]Message, error)
	Get(id uint) (*Message, error)

	GetMessagesByInboxId(inboxId uint) *[]Message
}
