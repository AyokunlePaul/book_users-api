package users

import (
	"fmt"
	"github.com/AyokunlePaul/book_users-api/datasources/mysql/users_db"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/utils/errors"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"log"
)

const (
	userInsertQuery = " INSERT INTO users (first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"
	userFetchQuery  = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?;"
)

func (user *User) GetUser() *response.BaseResponse {
	fetchStatement, fetchQueryParsingError := users_db.Client.Prepare(userFetchQuery)
	if fetchQueryParsingError != nil {
		log.Printf("internal server error: %s", fetchQueryParsingError)
		return response.NewInternalServerError(errors.UserFetchingError)
	}
	defer fetchStatement.Close()
	fetchResult := fetchStatement.QueryRow(user.Id)
	if scanError := fetchResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); scanError != nil {
		message := "User  doesn't exist in the database"
		return response.NewNotFoundError(message)
	}

	return nil
}

func (user *User) SaveUser() *response.BaseResponse {
	insertionStatement, insertionQueryParsingError := users_db.Client.Prepare(userInsertQuery)
	if insertionQueryParsingError != nil {
		log.Printf("internal server error: %s", insertionQueryParsingError)
		return response.NewInternalServerError(errors.UserCreationError)
	}
	defer insertionStatement.Close()

	insertionResult, insertionError := insertionStatement.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if insertionError != nil {
		log.Printf("internal server error: %s", insertionError)
		mySqlError, ok := insertionError.(*mysql.MySQLError)
		if !ok {
			return response.NewInternalServerError(errors.UserCreationError)
		}
		switch mySqlError.Number {
		case mysqlerr.ER_DUP_ENTRY:
			returnMessage := fmt.Sprintf("User with email %s %s", user.Email, errors.DuplicatedValueException)
			return response.NewBadRequestError(returnMessage)
		default:
			return response.NewInternalServerError(errors.UserCreationError)
		}
	}

	userId, lastIdError := insertionResult.LastInsertId()
	if lastIdError != nil {
		log.Printf("internal server error: %s", lastIdError)
		return response.NewInternalServerError(errors.UserCreationError)
	}
	user.Id = userId

	return nil
}
