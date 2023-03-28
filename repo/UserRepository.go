package repo

import (
	"ZepterUserService/model"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New() (*UserRepository, error) {
	repo := &UserRepository{}

	//TODO: Change this to env variables
	dsn := "host=userdb user=XML password=ftn dbname=XML_TEST port=5432 sslmode=disable"
	//dsn := "host=localhost user=XML password=ftn dbname=XML_TEST port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repo.db = db
	repo.db.AutoMigrate(&model.User{})

	return repo, nil
}

func (repo *UserRepository) Close() error {
	db, err := repo.db.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func (repo *UserRepository) SearchUsers(username string) []model.User {
	var users []model.User
	repo.db.Preload("Interests").Preload("Experiences").Model(&users).Where("LOWER(user_name) LIKE ?", "%"+strings.ToLower(username)+"%").Find(&users)
	return users
}

func (repo *UserRepository) GetByID(id int) model.User {
	var user model.User
	repo.db.Preload("Interests").Preload("Experiences").Model(&user).Find(&user, id)
	return user
}

func (repo *UserRepository) GetByUsername(username string) model.User {
	var user model.User
	repo.db.Preload("Interests").Preload("Experiences").Model(&user).Where("user_name  = ?", username).Find(&user)
	return user
}

func (repo *UserRepository) CreateUser(name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string, apiKey string) int {
	user := model.User{
		Name:        name,
		Email:       email,
		UserName:    username,
		Password:    password,
		Gender:      gender,
		PhoneNumber: phonenumber,
		DateOfBirth: dateofbirth,
		Biography:   biography,
		ApiKey:      apiKey,
	}

	if gender == "Male" || gender == "Female" {
		repo.db.Create(&user)
	} else {
		user.ID = 0
	}
	return int(user.ID)
}

func (repo *UserRepository) UpdateUser(id uint, name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string, isPrivate, changedPass bool) int {
	user := repo.GetByID(int(id))
	if changedPass {
		user.Forgotten = 0
	}
	user.Name = name
	user.Password = password
	user.Gender = gender
	user.PhoneNumber = phonenumber
	user.DateOfBirth = dateofbirth
	user.Biography = biography
	user.IsPrivate = isPrivate

	if gender == "Male" || gender == "Female" {
		repo.db.Save(&user)
	} else {
		user.ID = 0
	}
	return int(user.ID)
}

func (repo *UserRepository) Contains(elems []uint, v uint) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func (repo *UserRepository) Save(user model.User) {
	repo.db.Save(&user)
}
