package chat

import (
	"github.com/gorilla/websocket"
	"social-network/lib/repository"
)

type Chat struct {
	Repository        repository.MessageRepository
	PrivateDiscussion []PrivateChat
	GroupDiscussion   []GroupChat
}

type Peer struct {
	Id   int
	Name string
	Conn *websocket.Conn
}

type PrivateChat struct {
	Id   int
	User []Peer
}

type GroupChat struct {
	Id    int
	Name  string
	Users []Peer
}
