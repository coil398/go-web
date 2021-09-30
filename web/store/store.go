package store

import (
	"github.com/coil398/go-web/web/persistence"
	"github.com/coil398/go-web/web/repository"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	User repository.User
	Todo repository.Todo
}

func NewStore(db *sqlx.DB) *Store {
	user := persistence.NewUser(db)
	todo := persistence.NewTodo(db)
	return &Store{
		User: user,
		Todo: todo,
	}
}
