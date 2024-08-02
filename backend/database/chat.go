package database

import (
	"fmt"
	"social-network/lib/repository/models"
)

type MessageRepository struct{}

func (m *MessageRepository) NewMessage(senderId, receiverId int, text string) error {
	fmt.Println("This is a new message")
	return nil
}

func (m *MessageRepository) MessagePagination(userId int, page int, pageSize int) ([]models.Message, error) {
	fmt.Println("This is a message pagination")
	return []models.Message{}, nil
}
