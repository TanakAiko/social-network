package database

import (
	"fmt"
)

type PostRepository struct{}

func (post *PostRepository) NewPost(senderId, receiverId int, text string) error {
	fmt.Println("This is a new message")
	return nil
}
