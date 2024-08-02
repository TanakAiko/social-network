package main

import "social-network/services/chat"

func main() {
	chat := chat.NewChat()

	chat.Repository.NewMessage(1, 2, "Hello")
}
