package entity

import "time"

type Inbox struct {
	InboxId            uint `gorm:"primaryKey"` //`gorm:"column:inbox_id"`
	LastMessageContent string
	LastMessageDttm    time.Time
	LastMessageUser    uint
}

type InboxWithUserName struct {
	InboxId            uint
	LastMessageContent string
	LastMessageDttm    time.Time
	LastMessageUser    uint
	UserName           string
}

type InboxRepository interface {
	Create(*Inbox) error
	GetAll() (*[]Inbox, error)
	Get(id uint) (*Inbox, error)
	Update(*Inbox) error
	Delete(id uint) error

	GetAllInboxesByUsersIdx(usersIdx *[]uint) *[]Inbox
}
