package models

import (
	"azimio/database"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User defines the user in db
type User struct {
	gorm.Model
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Username         string `json:"username" gorm:"unique"`
	Phonenumber      string `json:"phonenumber"`
	Email            string `json:"email" gorm:"unique"`
	Password         string `json:"password"`
	Pollingstationid string `json:"pollingstationid" gorm:"type:text" gorm:"uniqueIndex"`
	Role             int    `json:"role" validate:"required, eq=ADMIN|eq=USER|eq=SADMIN"`
	Approved         bool   `json:"approved"`
}

type Polingdata struct {
	// gorm.Model
	Id               uint   `json:"id" gorm:"primary_key"`
	Pollingstationid string `json:"pollingstationid"`
	Candidate        string `json:"candidate" gorm:"type:text"`
	Scid             string `json:"scid" gorm:"type:text"`
	Ccode            string `json:"ccode" gorm:"type:text"`
	Cname            string `json:"cname" gorm:"type:text"`
	Scname           string `json:"scname" gorm:"type:text"`
	Pollingstation   string `json:"pollingstation"`
	Votes            uint   `json:"votes"`
	Registered       uint   `json:"registered"`
}

type Forms struct {
	Id               uint      `json:"id" gorm:"primary_key"`
	Pollingstationid string    `json:"pollingstationid"`
	Form             string    `json:"form"`
	Macaddress       string    `json:"macaddress"`
	Created_at       time.Time `json:"created_at" gorm:"column:created_at; type:timestamp; default: NOW(); not null; <-:create"`
	Updated_at       time.Time `json:"updated_at" gorm:"column:created_at; type:timestamp; default: NOW(); <-:update"`
	Username         string    `json:"userid"`
	Phonenumber      string    `json:"phonenumber"`
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	result := database.DBCon.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
