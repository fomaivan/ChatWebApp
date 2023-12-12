package service

import (
	"chat/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	userRepo entity.UserRepository
}

func NewUserService(userRepo entity.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Get(id uint) (*entity.User, error) {
	userDB, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return userDB, nil
	}
}

func comparePasswordWithHash(PasswordFromInput string, PasswordHashFromDB string) error {
	err := bcrypt.CompareHashAndPassword([]byte(PasswordHashFromDB), []byte(PasswordFromInput))
	return err
}

func generatePasswordHash(InputPassword string) (string, error) {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(InputPassword), bcrypt.DefaultCost)
	return string(PasswordHash), err
}

func (s *UserService) Update(user *entity.User) error {
	userDB, err := s.userRepo.Get(user.UserId)
	if err != nil {
		return err
	}
	var newUserPasswordHash string
	if comparePasswordWithHash(user.Password, userDB.Password) != nil {
		newUserPasswordHash, err = generatePasswordHash(user.Password)
		if err != nil {
			return err
		}
	}
	user.Password = newUserPasswordHash
	err = s.userRepo.Update(user)
	return err
}

func (s *UserService) Delete(user *entity.User) error {
	err := s.userRepo.Delete(user.UserId)
	return err
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UserService) Register(userReg *entity.UserRegister, chatService_ *entity.ChatService) error {
	PasswordHash, err := generatePasswordHash(userReg.Password)
	if err != nil {
		return err
	}
	userReg.Password = PasswordHash
	user := entity.User{
		UserRegister: *userReg,
	}
	user.CreateAt = time.Now()
	err = s.userRepo.Create(&user)
	if err != nil {
		return err
	}
	// Send first Message from me to new user!
	newUser, err := s.userRepo.GetByEmail(user.Email)
	firstMessageContent :=
		"Hello, " +
			newUser.UserName + ". " +
			"My name is Ivan and I'm the developer of this shit"

	(*chatService_).SendMessageFromTo(1, newUser.UserId, firstMessageContent)
	return err
}

func (s *UserService) Login(userLogin *entity.UserLogin) (uint, error) {
	User, err := s.userRepo.GetByEmail(userLogin.Email)
	if err != nil {
		return 0, err
	}
	err = comparePasswordWithHash(userLogin.Password, User.Password)
	if err != nil {
		return 0, err
	}
	return User.UserId, nil
}

func (s *UserService) GetAll() (*[]entity.User, error) {
	if users, err := s.userRepo.GetAll(); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (s *UserService) GetUsersByUserName(username string) (*[]entity.User, error) {
	return s.userRepo.GetUsersByUserName(username)
}

func (s *UserService) GetUserByUserName(username string) (*entity.User, error) {
	return s.userRepo.GetUserByUserName(username)
}
