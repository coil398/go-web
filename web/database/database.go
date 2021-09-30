package database

import (
	"fmt"
	"strings"

	"github.com/coil398/go-web/web/config"
	"github.com/jmoiron/sqlx"
)

func ConnectToDatabase(config *config.DBConfig) (*sqlx.DB, error) {
	loc := strings.Replace(config.TimeZone, "/", "%2F", 1)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", config.User, config.Password, config.Host, config.Port, config.Database, loc)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
