package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 					string `gorm:"primary_key;type:uuid" json:"id"`
	Username 			string `json:"username"`	
	Password 			string `json:"password"`
	Email 				string `json:"email"`
	CreatedAt 			time.Time `json:"created_at"`
	UpdatedAt 			time.Time `json:"updated_at"`
}

type LoginUser struct {
	Username 			string `json:"username"`
	Password 			string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Could not encrypt password")
	}
	return string(hash)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}