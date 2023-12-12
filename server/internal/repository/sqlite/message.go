package repo_sqlite

import (
	"chat/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type MessageSQLite struct {
	db *gorm.DB
}

func NewMessageSQLite(db *gorm.DB) *MessageSQLite {
	return &MessageSQLite{db: db}
}

func (r *MessageSQLite) Create(Message *entity.Message) error {
	if result := r.db.Create(Message); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MessageSQLite) Update(Message *entity.Message) error {
	result := r.db.Model(Message).Updates(Message)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MessageSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Message{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MessageSQLite) GetAll() (*[]entity.Message, error) {
	var Messages []entity.Message
	if result := r.db.Find(&Messages); result.Error != nil {
		return nil, result.Error
	} else {
		return &Messages, nil
	}
}

func (r *MessageSQLite) Get(id uint) (*entity.Message, error) {
	var Message entity.Message
	if result := r.db.Where("id = ?", id).First(&Message); result.Error == nil {
		return &Message, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &Message, entity.ErrMessageNotFound
	} else {
		return &Message, result.Error
	}
}

func (r *MessageSQLite) GetMessagesByInboxId(inboxId uint) *[]entity.Message {
	var Messages []entity.Message
	r.db.Table("messages").Where("inbox_id = ?", inboxId).Find(&Messages)
	return &Messages
}
