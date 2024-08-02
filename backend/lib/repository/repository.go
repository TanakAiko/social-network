package repository

import (
	"social-network/lib/repository/models"
)

type MessageRepository interface {
	NewMessage(senderId, receiverId int, text string) error
	MessagePagination(userId int, page int, pageSize int) ([]models.Message, error)
}


type UserRepository interface {
	NewUser() error
	GetUser(id int) error
}