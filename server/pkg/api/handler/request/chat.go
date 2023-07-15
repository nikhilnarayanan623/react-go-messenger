package request

type Chat struct {
	OtherUserID uint `json:"other_user_id" binding:"required"`
}

type Message struct {
	Content string `json:"content" binding:"required"`
}
