package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"task_system_go/database"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" autoIncrement:"true"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `gorm:"index:idx_username,unique" json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (user *User) CreateUser() error {
	var alreadyUser User
	res := database.Database.Where("username = ?", user.Username).First(&alreadyUser)
	if res.RowsAffected == 1 {
		return errors.New("user already exists")
	}
	result := database.Database.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) ValidatePassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	return err == nil
}
