package model

type User struct {
	CreatedAt string  `json:"createdAt" db:"created_at"`
	Email     string  `json:"email" db:"email"`
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Note      *string `json:"note" db:"note"`
	Password  string  `json:"password" db:"password"`
	UpdatedAt string  `json:"updatedAt" db:"updated_at"`
}
