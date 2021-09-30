package repository

import "github.com/coil398/go-web/web/model"

type User interface {
	PostUser(user *model.User) (*model.User, error)
	GetUser(id int) (*model.User, error)
	PatchUser(user *model.User) (*model.User, error)
}

type Todo interface {
	GetTodos(userId int) ([]*model.Todo, error)
	PostTodo(userId int, todo *model.Todo) (*model.Todo, error)
	PatchTodo(userId int, todo *model.Todo) (*model.Todo, error)
}
