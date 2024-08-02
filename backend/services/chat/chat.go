package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"social-network/database"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewChat() *Chat {
	return &Chat{
		Repository:        &database.MessageRepository{},
		PrivateDiscussion: []PrivateChat{},
		GroupDiscussion:   []GroupChat{},
	}
}

func (c *Chat) AddPrivateChat(private PrivateChat) {
	c.PrivateDiscussion = append(c.PrivateDiscussion, private)
}

func (c *Chat) AddGroupChat(group GroupChat) {
	c.GroupDiscussion = append(c.GroupDiscussion, group)
}


