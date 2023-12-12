package entity

type ChatService interface {
	SendMessageFromTo(fromId uint, toId uint, messageContent string) (*Message, error)
	GetAllInboxesByUserId(userId uint) *[]Inbox
	GetMessagesByUserAndContact(userId uint, contactId uint) *[]Message
	GetMessagesByInboxId(inboxId uint) *[]Message
	GetAllInboxesWithUserNameByUserId(userId uint) *[]InboxWithUserName
	//GetInboxWithUserName(userId uint) *InboxWithUserName
	GetInboxIdByTwoUsers(users *TwoUserIdx) (uint, error)
	GetInboxByInboxId(inboxId uint) (*Inbox, error)
	CreateInboxAndInboxUser(firstId uint, secondId uint) (*Inbox, error)
}
