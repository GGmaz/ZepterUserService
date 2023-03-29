package service

import (
	"math/rand"
	"zepter/model"
	"zepter/repo"

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

func (s *UserService) SearchUsers(country string, page, limit int) []model.User {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}
	users := s.userRepo.SearchUsers(country, page, limit)

	return users
}

func (s *UserService) GetByID(id int) model.User {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetByUsername(username string) model.User {
	return s.userRepo.GetByUsername(username)
}

func (s *UserService) CreateUser(firstName, email, password, username, lastName, country string) int {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return s.userRepo.CreateUser(firstName, email, string(hashedPassword), username, lastName, country)
}

func (s *UserService) UpdateUser(id uint, firstName string, lastName string, country string, password string) int {
	if password != s.GetByID(int(id)).Password {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
		return s.userRepo.UpdateUser(id, firstName, lastName, country, string(hashedPassword))
	}
	return s.userRepo.UpdateUser(id, firstName, lastName, country, password)
}

func (s *UserService) DeleteUser(id int) {
	s.userRepo.Delete(id)
	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
