package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderUserID   uint   `gorm:"not null" json:"sender_user_id"`
	ReceiverUserID uint   `gorm:"not null" json:"receiver_user_id"`
	Message        string `gorm:"type:text" json:"message"`
	IsRead         bool   `gorm:"default:true" json:"is_read"`
}