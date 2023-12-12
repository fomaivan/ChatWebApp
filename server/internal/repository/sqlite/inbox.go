package repo_sqlite

import (
	"chat/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type InboxSQLite struct {
	db *gorm.DB
}

func NewInboxSQLite(db *gorm.DB) *InboxSQLite {
	return &InboxSQLite{db: db}
}

func (r *InboxSQLite) GetAll() (*[]entity.Inbox, error) {
	var Inboxes []entity.Inbox
	if result := r.db.Find(&Inboxes); result.Error != nil {
		return nil, result.Error
	} else {
		return &Inboxes, nil
	}
}

func (r *InboxSQLite) Get(id uint) (*entity.Inbox, error) {
	var Inbox entity.Inbox
	if result := r.db.Where("inbox_id = ?", id).First(&Inbox); result.Error == nil {
		return &Inbox, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &Inbox, entity.ErrInboxNotFound
	} else {
		return &Inbox, result.Error
	}
}

func (r *InboxSQLite) Create(Inbox *entity.Inbox) error {
	if result := r.db.Create(Inbox); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxSQLite) Update(Inbox *entity.Inbox) error {
	result := r.db.Model(Inbox).Updates(Inbox)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Inbox{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxSQLite) GetAllInboxesByUsersIdx(usersIdx *[]uint) *[]entity.Inbox {
	var Inboxes []entity.Inbox
	r.db.Order("last_message_dttm desc").Find(&Inboxes, usersIdx)
	return &Inboxes
}
