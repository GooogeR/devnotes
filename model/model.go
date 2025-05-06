package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Note struct {
	ID      string `json:"id"`
	OwnerID string `json:"owner_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
