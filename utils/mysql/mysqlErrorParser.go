package mysql

import (
	"database/sql"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/go-sql-driver/mysql"
)

func HandleError(rawError error, optionalValue ...interface{}) *response.BaseResponse {
	mysqlError, ok := rawError.(*mysql.MySQLError)
	if !ok {
		switch rawError {
		case sql.ErrNoRows:

		}
	} else {
		switch mysqlError.Number {

		}
	}
	return nil
}
