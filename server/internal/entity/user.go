package entity

import (
	_ "image/jpeg"
	_ "image/png"
	"time"
)

//import "image/png"

type User struct {
	UserId   uint `gorm:"primaryKey"`
	CreateAt time.Time

	UserRegister
}

type UserLogin struct {
	Email    string `gorm:"unique"`
	Password string
}

type UserRegister struct {
	UserLogin
	UserName string `gorm:"unique"`
	//UserImage image.Image
}

type TwoUserIdx struct {
	FirstUser  uint
	SecondUser uint
}

type UserRepository interface {
	Create(*User) error
	GetAll() (*[]User, error)
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(*User) error
	Delete(id uint) error
	GetUserNames(idxArr *[]uint) *[]User
	GetUsersByUserName(username string) (*[]User, error)
	GetUserByUserName(username string) (*User, error)
}

type UserService interface {
	Get(id uint) (*User, error)
	Update(user *User) error
	Delete(user *User) error
	GetAll() (*[]User, error)
	GetUserByEmail(email string) (*User, error)
	GetUsersByUserName(username string) (*[]User, error)
	GetUserByUserName(username string) (*User, error)

	Register(userReg *UserRegister, chatService_ *ChatService) error
	Login(userLogin *UserLogin) (uint, error)
}
