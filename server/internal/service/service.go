package service

import (
	"chat/internal/entity"
	"chat/internal/repository"
)

type Service struct {
	User entity.UserService
	Chat entity.ChatService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
		Chat: NewChatService(repo.Message, repo.Inbox, repo.InboxUser),
	}
}
