package entity

type InboxUser struct {
	Id      uint `gorm:"primaryKey"`
	InboxId uint
	UserId  uint
}

type InboxUserRepository interface {
	Create(*InboxUser) error
	GetAll() (*[]InboxUser, error)
	Get(id uint) (*InboxUser, error)
	Update(*InboxUser) error
	Delete(id uint) error

	GetInboxIdByTwoUsers(firstId uint, secondId uint) (uint, error)
	GetAllInboxUsersIdxByUser(userId uint) *[]uint
	GetUsersIdxByInbox(inboxIdx uint) *[]uint
	GetAllContactsByUserId(userId uint) *[]uint
	GetAllInboxesWithUserNameByUserId(userId uint) *[]InboxWithUserName
	//GetContactIdByUserIdAndInboxId(userId uint, inboxId uint) (uint, error)
}
