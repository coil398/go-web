package persistence

import (
	"github.com/coil398/go-web/web/model"
	"github.com/coil398/go-web/web/repository"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{
		db,
	}
}

func (u *User) GetUser(id int) (*model.User, error) {
	stmt, err := u.db.Preparex(`SELECT * FROM users WHERE users.id = ?`)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowx(id)
	var user model.User
	if err := row.StructScan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) PostUser(user *model.User) (*model.User, error) {
	stmt, err := u.db.PrepareNamed(`INSERT INTO users (name, email, password, note) VALUES (:name, :email, :password, :note)`)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(user)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = int(id)
	return user, nil
}

func (u *User) PatchUser(user *model.User) (*model.User, error) {
	stmt, err := u.db.PrepareNamed(`UPDATE users SET name = COALESCE(:name, name) AND email = COALESCE(:email, email) AND password = COALESCE(:password, password) AND note = COALESCE(:note, note)`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(user); err != nil {
		return nil, err
	}
	return user, nil
}
