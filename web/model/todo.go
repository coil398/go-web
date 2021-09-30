package model

type Todo struct {
	Content   string  `json:"content" db:"content"`
	CreatedAt string  `json:"createdAt" db:"created_at"`
	Done      bool    `json:"done" db:"done"`
	Id        int     `json:"id" db:"id"`
	Note      *string `json:"note" db:"note"`
	UpdatedAt string  `json:"updatedAt" db:"updated_at"`
	UserId    int     `json:"userId" db:"user_id"`
}
