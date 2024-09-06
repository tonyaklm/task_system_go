package models

import (
	"errors"
	"gorm.io/gorm"
	"task_system_go/database"
)

type Post struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" autoIncrement:"true"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Title   string `gorm:"not null" default:"" json:"title" binding:"required"`
	Content string `gorm:"not null" default:"" json:"content" binding:"required"`
}

func (post *Post) CreatePost() error {
	result := database.Database.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (post *Post) FindById(postId uint) error {
	res := database.Database.Where("id = ?", postId).First(&post)
	if res.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}
