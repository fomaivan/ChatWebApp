package repository

import (
	"chat/internal/entity"
	repo_sqlite "chat/internal/repository/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	User      entity.UserRepository
	Message   entity.MessageRepository
	Inbox     entity.InboxRepository
	InboxUser entity.InboxUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:      repo_sqlite.NewUserSQLite(db),
		Message:   repo_sqlite.NewMessageSQLite(db),
		Inbox:     repo_sqlite.NewInboxSQLite(db),
		InboxUser: repo_sqlite.NewInboxUserSQLite(db),
	}
}
