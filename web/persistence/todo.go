package persistence

import (
	"github.com/coil398/go-web/web/model"
	"github.com/coil398/go-web/web/repository"
	"github.com/jmoiron/sqlx"
)

type Todo struct {
	db *sqlx.DB
}

func NewTodo(db *sqlx.DB) repository.Todo {
	return &Todo{
		db,
	}
}

func (t *Todo) GetTodos(userId int) ([]*model.Todo, error) {
	stmt, err := t.db.Preparex(`SELECT * FROM todos WHERE todos.user_id = ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Queryx(userId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo

	for rows.Next() {
		var t model.Todo
		if err := rows.StructScan(&t); err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}
	return todos, nil
}

func (t *Todo) PostTodo(userId int, todo *model.Todo) (*model.Todo, error) {
	stmt, err := t.db.PrepareNamed(`INSERT INTO todos (content, note) VALUES (:content, :note)`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Todo) PatchTodo(userId int, todo *model.Todo) (*model.Todo, error) {
	stmt, err := t.db.PrepareNamed(`UPDATE todos SET content = COALESCE(:content, content) AND note = COALESCE(:note, note) AND done = COALESCE(:done, done) WHERE todos.user_id = :user_id`)
	if err != nil {
		return nil, err
	}
	if _, err := stmt.Exec(todo); err != nil {
		return nil, err
	}
	return todo, nil
}
