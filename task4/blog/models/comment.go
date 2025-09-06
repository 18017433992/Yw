package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    User `gorm:"foreignkey:UserID"`
	PostID  uint
	Post    Post `gorm:"PostID"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}
