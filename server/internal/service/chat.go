package service

import (
	"chat/internal/entity"
	"errors"
	"time"
)

type ChatService struct {
	MessageRepo   entity.MessageRepository
	InboxRepo     entity.InboxRepository
	InboxUserRepo entity.InboxUserRepository
}

func NewChatService(
	MessageRepo entity.MessageRepository,
	InboxRepo entity.InboxRepository,
	InboxUserRepo entity.InboxUserRepository) *ChatService {
	return &ChatService{
		MessageRepo:   MessageRepo,
		InboxRepo:     InboxRepo,
		InboxUserRepo: InboxUserRepo,
	}
}

func (s *ChatService) SendMessageFromTo(fromId uint, toId uint, messageContent string) (*entity.Message, error) {
	SendMessageTime := time.Now()
	InboxIdx, err := s.InboxUserRepo.GetInboxIdByTwoUsers(fromId, toId)

	// There is no such dialogue -> create a dialogue and send message.
	if errors.Is(err, entity.ErrInboxUserNotFound) {
		// create Inbox
		NewInbox := entity.Inbox{
			LastMessageContent: messageContent,
			LastMessageDttm:    SendMessageTime,
			LastMessageUser:    fromId,
		}
		s.InboxRepo.Create(&NewInbox)
		InboxIdx = NewInbox.InboxId

		// create InboxUser
		NewInboxUser1 := entity.InboxUser{InboxId: NewInbox.InboxId, UserId: fromId}
		NewInboxUser2 := entity.InboxUser{InboxId: NewInbox.InboxId, UserId: toId}
		s.InboxUserRepo.Create(&NewInboxUser1)
		s.InboxUserRepo.Create(&NewInboxUser2)

		// create Message
		NewMessage := entity.Message{
			InboxId:  InboxIdx,
			UserId:   fromId,
			Content:  messageContent,
			SendTime: SendMessageTime,
		}
		s.MessageRepo.Create(&NewMessage)
		return &NewMessage, nil
	} else if err == nil {
		// create Message
		NewMessage := entity.Message{
			InboxId:  InboxIdx,
			UserId:   fromId,
			Content:  messageContent,
			SendTime: SendMessageTime,
		}
		s.MessageRepo.Create(&NewMessage)

		// Update LastMessage in Inbox
		s.InboxRepo.Update(&entity.Inbox{
			InboxId:            InboxIdx,
			LastMessageContent: messageContent,
			LastMessageDttm:    SendMessageTime,
			LastMessageUser:    fromId,
		})
		return &NewMessage, nil
	} else {
		// Error
		return nil, err
	}
}

func (s *ChatService) GetAllInboxesByUserId(userId uint) *[]entity.Inbox {
	InboxesIdx := s.InboxUserRepo.GetAllInboxUsersIdxByUser(userId)
	return s.InboxRepo.GetAllInboxesByUsersIdx(InboxesIdx)
}

func (s *ChatService) GetAllInboxesWithUserNameByUserId(userId uint) *[]entity.InboxWithUserName {
	return s.InboxUserRepo.GetAllInboxesWithUserNameByUserId(userId)
}

func (s *ChatService) GetMessagesByUserAndContact(userId uint, contactId uint) *[]entity.Message {
	if InboxId, err := s.InboxUserRepo.GetInboxIdByTwoUsers(userId, contactId); err != nil {
		return nil
	} else {
		return s.MessageRepo.GetMessagesByInboxId(InboxId)
	}
}

func (s *ChatService) GetMessagesByInboxId(inboxId uint) *[]entity.Message {
	return s.MessageRepo.GetMessagesByInboxId(inboxId)
}

func (s *ChatService) GetInboxIdByTwoUsers(users *entity.TwoUserIdx) (uint, error) {
	return s.InboxUserRepo.GetInboxIdByTwoUsers(users.FirstUser, users.SecondUser)
}

func (s *ChatService) GetInboxByInboxId(inboxId uint) (*entity.Inbox, error) {
	return s.InboxRepo.Get(inboxId)
}

func (s *ChatService) CreateInboxAndInboxUser(firstId uint, secondId uint) (*entity.Inbox, error) {
	SendMessageTime := time.Now()
	NewInbox := entity.Inbox{
		LastMessageContent: "",
		LastMessageDttm:    SendMessageTime,
		LastMessageUser:    firstId,
	}
	err := s.InboxRepo.Create(&NewInbox)
	if err != nil {
		return nil, err
	}
	NewInboxUser1 := entity.InboxUser{InboxId: NewInbox.InboxId, UserId: firstId}
	NewInboxUser2 := entity.InboxUser{InboxId: NewInbox.InboxId, UserId: secondId}
	err = s.InboxUserRepo.Create(&NewInboxUser1)
	if err != nil {
		return nil, err
	}
	err = s.InboxUserRepo.Create(&NewInboxUser2)
	if err != nil {
		return nil, err
	}
	return &NewInbox, nil
}
