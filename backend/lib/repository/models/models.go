package models

type Message struct {
	Id         int    `json:"id"`
	Message    string `json:"content"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
}

type User struct {
}

type Post struct {
	Id         int      `json:"id"`
	UserId     int      `json:"user_id"`
	Nickname   string   `json:"nickname"`
	Categorie  []string `json:"categorie"`
	LikedBy    []string `json:"likedBy"`
	DisLikedBy []string `json:"dislikedBy"`
	Content    string   `json:"content"`
	Img        string   `json:"img"`
	NbrLike    int      `json:"nbrLike"`
	NbrDislike int      `json:"nbrDislike"`
	CreatedAt  string   `json:"created_at"`
}
