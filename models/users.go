package models

import (
	"errors"

	"github.com/jinzhu/gorm"

	// imported for the effects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is a custom error we return when a resource we are looking for is not present in the db
	ErrNotFound = errors.New("models: resource not found")
)

// User is the database model for our customer
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

// UserService stores the information required for abstracting functions related to the db
type UserService struct {
	db *gorm.DB
}

// NewUserService is an abstraction layer providing us a connection with the db
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// Close is a function that is used to close the connection with the db
func (us *UserService) Close() error {
	return us.db.Close()
}

// ByID is used to search a user by ID from the db
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DestructiveReset drops the user table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// Create is used to add a new user
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return nil
}

// ByEmail is used to search a user by email from the db
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
