package main

import (
	"net/http"

	"github.com/coil398/go-web/web/Openapi"
	"github.com/coil398/go-web/web/model"
	"github.com/coil398/go-web/web/store"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type server struct {
	store *store.Store
}

func (s *server) PostUser(ctx echo.Context) error {
	ctx.Logger().Info("PostUser")
	u := Openapi.PostUserJSONRequestBody{}
	if err := ctx.Bind(&u); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := &model.User{
		Name:     u.Name,
		Email:    u.Email,
		Note:     u.Note,
		Password: string(hashedPassword),
	}

	s.store.User.PostUser(user)

	return ctx.JSON(http.StatusOK, u)
}

func (s *server) GetUsersUserId(ctx echo.Context, userId int) error {
	ctx.Logger().Info("GetUsersUserId")
	user, err := s.store.User.GetUser(userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (s *server) PatchUsersUserId(ctx echo.Context, userId int) error {
	ctx.Logger().Info("PatchUsersUserId")
	return nil
}

func (s *server) PostUsersUserIdTodo(ctx echo.Context, userId string) error {
	ctx.Logger().Info("PostUsersUserIdTodo")
	return nil
}

func (s *server) PatchUsersUserIdTodoTodoId(ctx echo.Context, userId string, todoId string) error {
	ctx.Logger().Info("PatchUsersUserIdTodoTodoId")
	return nil
}

func (s *server) GetUsersUserIdTodos(ctx echo.Context, userId string) error {
	ctx.Logger().Info("GetUsersUserIdTodos")
	return nil
}
