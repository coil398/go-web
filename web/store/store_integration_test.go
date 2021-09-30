//go:build integration

package store

import (
	"os"
	"testing"

	"github.com/coil398/go-web/web/config"
	"github.com/coil398/go-web/web/database"
	"github.com/coil398/go-web/web/model"
	"github.com/stretchr/testify/assert"
)

var store *Store

func dbHandlingWrapper(m *testing.M) int {
	c, err := config.ReadDBConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectToDatabase(c)
	if err != nil {
		panic(err)
	}
	store = NewStore(db)
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(dbHandlingWrapper(m))
}

func TestPostUser(t *testing.T) {
	user := &model.User{
		Name:     "user",
		Email:    "foo@bar.baz",
		Password: "hashedpassword",
	}

	res, err := store.User.PostUser(user)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.Name, res.Name)
	assert.Equal(t, user.Email, res.Email)
	assert.Equal(t, user.Password, res.Password)
}

func TestGetUser(t *testing.T) {
	user := &model.User{
		Name:     "user",
		Email:    "hoge@fuga.piyo",
		Password: "hashedpassword",
	}

	res, err := store.User.PostUser(user)
	if err != nil {
		t.Fatal(err)
	}

	got, err := store.User.GetUser(res.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.Name, got.Name)
	assert.Equal(t, user.Email, got.Email)
	assert.Equal(t, user.Password, got.Password)
}

func TestPostDuplicatedEmailUser_Error(t *testing.T) {
	user := &model.User{
		Name:     "user",
		Email:    "hoge@fuga.piyo",
		Password: "hashedpassword",
	}
	another := &model.User{
		Name:     "another",
		Email:    "hoge@fuga.piyo",
		Password: "hashedpassword",
	}
	store.User.PostUser(user)
	_, err := store.User.PostUser(another)
	if err == nil {
		t.Fatalf("PostDuplicatedEmailUser must be failed")
	}
}
