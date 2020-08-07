package users_db

import (
	"database/sql"
	"fmt"
	"github.com/AyokunlePaul/book_users-api/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() { //&parseTime=true
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var dbError error
	Client, dbError = sql.Open("mysql", dataSourceName)
	if dbError != nil {
		panic(dbError)
	}

	if dbError = Client.Ping(); dbError != nil {
		panic(dbError)
	}

	logger.Info("User database successfully configured!")
}
