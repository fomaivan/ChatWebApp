package repo_sqlite

import (
	"chat/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type InboxUserSQLite struct {
	db *gorm.DB
}

func NewInboxUserSQLite(db *gorm.DB) *InboxUserSQLite {
	return &InboxUserSQLite{db: db}
}

func (r *InboxUserSQLite) GetAll() (*[]entity.InboxUser, error) {
	var InboxUsers []entity.InboxUser
	if result := r.db.Find(&InboxUsers); result.Error != nil {
		return nil, result.Error
	} else {
		return &InboxUsers, nil
	}
}

func (r *InboxUserSQLite) Get(id uint) (*entity.InboxUser, error) {
	var InboxUser entity.InboxUser
	if result := r.db.Where("id = ?", id).First(&InboxUser); result.Error == nil {
		return &InboxUser, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &InboxUser, entity.ErrInboxUserNotFound
	} else {
		return &InboxUser, result.Error
	}
}

func (r *InboxUserSQLite) Create(InboxUser *entity.InboxUser) error {
	if result := r.db.Create(InboxUser); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxUserSQLite) Update(InboxUser *entity.InboxUser) error {
	result := r.db.Model(InboxUser).Updates(InboxUser)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxUserSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.InboxUser{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *InboxUserSQLite) GetInboxIdByTwoUsers(firstId uint, secondId uint) (uint, error) {
	var InboxIdx []uint
	r.db.Raw(
		"SELECT first.inbox_id "+
			"FROM inbox_users first "+
			"INNER JOIN "+
			"(SELECT * FROM inbox_users WHERE user_id = ?) second ON "+
			"first.inbox_id = second.inbox_id "+
			"WHERE first.user_id = ?;",
		firstId, secondId,
	).Scan(&InboxIdx)
	if len(InboxIdx) == 0 {
		return 1, entity.ErrInboxUserNotFound
	} else {
		return InboxIdx[0], nil
	}
}

func (r *InboxUserSQLite) GetAllInboxUsersIdxByUser(userId uint) *[]uint {
	var InboxesIdx []uint
	r.db.Raw(
		"SELECT inbox_id "+
			"FROM inbox_users "+
			"WHERE user_id = ?", userId,
	).Scan(&InboxesIdx)
	return &InboxesIdx
}

func (r *InboxUserSQLite) GetAllContactsByUserId(userId uint) *[]uint {
	var UsersIdx []uint
	InboxIdx := r.GetAllInboxUsersIdxByUser(userId)
	r.db.Table("inbox_users").
		Select("user_id").
		Where("user_id <> ? AND inbox_id IN ?", userId, InboxIdx).Find(&UsersIdx)
	return &UsersIdx
}

func (r *InboxUserSQLite) GetUsersIdxByInbox(inboxIdx uint) *[]uint {
	var users []uint
	r.db.Raw(
		"SELECT user_id "+
			"FROM inbox_users "+
			"WHERE inbox_id = ?", inboxIdx,
	).Scan(&users)
	return &users
}

func (r *InboxUserSQLite) GetAllInboxesWithUserNameByUserId(userId uint) *[]entity.InboxWithUserName {
	var result *[]entity.InboxWithUserName
	InboxesIdx := r.GetAllInboxUsersIdxByUser(userId)
	//fmt.Println(len(*InboxesIdx))
	r.db.Raw(
		"SELECT inb.inbox_id, inb.last_message_content, "+
			"inb.last_message_dttm, inb.last_message_user, "+
			"users.user_name "+
			"FROM inbox_users in_u "+
			"LEFT JOIN inboxes inb ON inb.inbox_id = in_u.inbox_id  "+
			"LEFT JOIN users ON users.user_id = in_u.user_id "+
			"WHERE in_u.user_id <> ? AND in_u.inbox_id IN ? "+
			"ORDER BY inb.last_message_dttm DESC;", userId, *InboxesIdx).Scan(&result)
	return result
}

//func (r *InboxUserSQLite) GetContactIdByUserIdAndInboxId(userId uint, inboxId uint) (uint, error) {
//	var ContactId uint
//	err := r.db.Table("inbox_users").
//		Select("user_id").
//		Where("user_id <> ? AND inbox_id = ?", userId, inboxId).Find(&ContactId)
//	return ContactId, err.Error
//}
