package repo_sqlite

import (
	"chat/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQliteDB(dbUri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Inbox{},
		&entity.InboxUser{},
		&entity.Message{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
