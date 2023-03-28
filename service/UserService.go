package service

import (
	"ZepterUserService/model"
	"ZepterUserService/repo"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func New() (*UserService, error) {

	userRepo, err := repo.New()
	if err != nil {
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "NewUserService"}).Error("Error creating User Repository.")
		return nil, err
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "NewUserService"}).Info("Successfully created User Service.")
	return &UserService{
		userRepo: userRepo,
	}, nil
}

func (s *UserService) SearchUsers(username string, id uint) []model.User {
	users := s.userRepo.SearchUsers(username)
	if id == 0 {
		return users
	}
	blockedIds := s.userRepo.FindBlockedForUserId(id)
	i := 0
	for _, user := range users {
		if !s.userRepo.Contains(blockedIds, user.ID) && user.ID != id {
			users[i] = user
			i++
		}
	}
	users = users[:i]
	return users
}

func (s *UserService) GetByID(id int) model.User {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetByUsername(username string) model.User {
	return s.userRepo.GetByUsername(username)
}

func (s *UserService) CreateUser(name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string) int {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return s.userRepo.CreateUser(name, email, string(hashedPassword), username, gender, phonenumber, dateofbirth, biography, string(apiKey))
}

func (s *UserService) UpdateUser(id uint, name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string, isPrivate bool) int {
	if password != s.GetByID(int(id)).Password {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
		return s.userRepo.UpdateUser(id, name, email, string(hashedPassword), username, gender, phonenumber, dateofbirth, biography, isPrivate, true)
	}
	return s.userRepo.UpdateUser(id, name, email, password, username, gender, phonenumber, dateofbirth, biography, isPrivate, false)
}

func (s *UserService) ForgotPassword(username string) int {
	user := s.GetByUsername(username)
	if user.ID == 0 {
		return 0
	}
	newPass := GenerateRandomString(10)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPass), 8)
	user.Password = string(hashedPassword)
	user.Forgotten = 1
	s.userRepo.Save(user)

	SendActivationMail(user.Email, newPass, "")
	return 1
}
