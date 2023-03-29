package repo

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
	"time"
	"zepter/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New() (*UserRepository, error) {
	repo := &UserRepository{}

	godotenv.Load(".env")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
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
	repo.db.Model(&users).Where("LOWER(country) LIKE ?", "%"+strings.ToLower(username)+"%").Find(&users)
	return users
}

func (repo *UserRepository) GetByID(id int) model.User {
	var user model.User
	repo.db.Model(&user).Find(&user, id)
	return user
}

func (repo *UserRepository) GetByUsername(username string) model.User {
	var user model.User
	repo.db.Model(&user).Where("username  = ?", username).Find(&user)
	return user
}

func (repo *UserRepository) CreateUser(firstName, email, password, username, lastName, country string) int {
	user := model.User{
		FirstName: firstName,
		Email:     email,
		LastName:  lastName,
		Country:   country,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	repo.db.Create(&user)

	return int(user.ID)
}

func (repo *UserRepository) UpdateUser(id uint, firstName string, lastName string, country string, password string) int {
	user := repo.GetByID(int(id))

	user.FirstName = firstName
	user.LastName = lastName
	user.Password = password
	user.Country = country
	user.UpdatedAt = time.Now()

	repo.db.Save(&user)

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

func (repo *UserRepository) Delete(id int) {
	repo.db.Delete(&model.User{}, id)
}
