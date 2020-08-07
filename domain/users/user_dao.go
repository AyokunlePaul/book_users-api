package users

import (
	"database/sql"
	"fmt"
	"github.com/AyokunlePaul/book_users-api/datasources/mysql/users_db"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/logger"
	"github.com/AyokunlePaul/book_users-api/utils/errors"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"log"
)

const (
	userInsertQuery       = "INSERT INTO users (first_name, last_name, email, created_at, updated_at, status, password) VALUES (?, ?, ?, ?, ?, ?, ?);"
	userFetchQuery        = "SELECT * FROM users WHERE id=?;"
	userFetchAllQuery     = "SELECT * FROM users;"
	userUpdateQuery       = "UPDATE users SET first_name=?, last_name=?, email=?, updated_at=? WHERE id=?;"
	userDeleteQuery       = "DELETE FROM users WHERE id=?"
	userFindByStatusQuery = "SELECT * FROM users WHERE status=?;"
)

func (user *User) Get() *response.BaseResponse {
	fetchStatement, fetchQueryParsingError := users_db.Client.Prepare(userFetchQuery)
	if fetchQueryParsingError != nil {
		logger.Error(errors.ZapUserFetchError, fetchQueryParsingError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer fetchStatement.Close()
	fetchResult := fetchStatement.QueryRow(user.Id)
	if scanError := fetchResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status, &user.Password); scanError != nil {
		logger.Error(errors.ZapUserFetchError, scanError)
		switch scanError {
		case sql.ErrNoRows:
			message := "no user with the given id"
			return response.NewNotFoundError(message)
		default:
			message := "error fetching the user"
			return response.NewNotFoundError(message)
		}
	}
	return nil
}

func (user *User) Save() *response.BaseResponse {
	insertStatement, insertQueryParsingError := users_db.Client.Prepare(userInsertQuery)
	if insertQueryParsingError != nil {
		logger.Error(errors.ZapUserCreateError, insertQueryParsingError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer insertStatement.Close()

	insertResult, insertError := insertStatement.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt, user.Status, user.Password)
	if insertError != nil {
		logger.Error(errors.ZapUserCreateError, insertError)
		mySqlError, ok := insertError.(*mysql.MySQLError)
		if !ok {
			return response.NewInternalServerError(errors.UserCreationError)
		}
		switch mySqlError.Number {
		case mysqlerr.ER_DUP_ENTRY:
			returnMessage := fmt.Sprintf("%s %s", user.Email, "already used")
			return response.NewBadRequestError(returnMessage)
		default:
			return response.NewInternalServerError(errors.InternalServerErrorMessage)
		}
	}

	userId, lastIdError := insertResult.LastInsertId()
	if lastIdError != nil {
		logger.Error(errors.ZapUserCreateError, lastIdError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *response.BaseResponse {
	updateStatement, updateQueryParseError := users_db.Client.Prepare(userUpdateQuery)
	if updateQueryParseError != nil {
		logger.Error(errors.ZapUserUpdateError, updateQueryParseError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer updateStatement.Close()

	_, updateError := updateStatement.Exec(user.FirstName, user.LastName, user.Email, user.UpdatedAt, user.Id)
	if updateError != nil {
		logger.Error(errors.ZapUserUpdateError, updateError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}

	return nil
}

func (user *User) Delete() *response.BaseResponse {
	deleteStatement, deleteQueryParseError := users_db.Client.Prepare(userDeleteQuery)
	if deleteQueryParseError != nil {
		logger.Error(errors.ZapUserDeleteError, deleteQueryParseError)
		return response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer deleteStatement.Close()

	_, deleteError := deleteStatement.Exec(user.Id)
	if deleteError != nil {
		logger.Error(errors.ZapUserDeleteError, deleteError)
		return response.NewInternalServerError("error updating the user")
	}

	return nil
}

func GetAll() ([]User, *response.BaseResponse) {
	var users = make([]User, 0)
	fetchAllStatement, fetchAllQueryParsingError := users_db.Client.Prepare(userFetchAllQuery)
	if fetchAllQueryParsingError != nil {
		logger.Error(errors.ZapUserFetchError, fetchAllQueryParsingError)
		return nil, response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer fetchAllStatement.Close()

	fetchAllResult, fetchAllError := fetchAllStatement.Query()
	if fetchAllError != nil {
		logger.Error(errors.ZapUserFetchError, fetchAllError)
		return nil, response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer log.Fatal(fetchAllResult.Close())

	for fetchAllResult.Next() {
		user := User{}
		scanError := fetchAllResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if scanError != nil {
			logger.Error(errors.ZapUserFetchError, scanError)
			return users, nil
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUsersByStatus(status string) ([]User, *response.BaseResponse) {
	var users = make([]User, 0)
	fetchByStatusStatement, fetchByStatusQueryParseError := users_db.Client.Prepare(userFindByStatusQuery)
	if fetchByStatusQueryParseError != nil {
		logger.Error(errors.ZapUserFetchError, fetchByStatusQueryParseError)
		return nil, response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer fetchByStatusStatement.Close()

	fetchByStatusResult, fetchByStatusError := fetchByStatusStatement.Query(status)
	if fetchByStatusError != nil {
		logger.Error(errors.ZapUserFetchError, fetchByStatusError)
		return nil, response.NewInternalServerError(errors.InternalServerErrorMessage)
	}
	defer fetchByStatusResult.Close()

	for fetchByStatusResult.Next() {
		user := User{}
		if scanError := fetchByStatusResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status, &user.Password); scanError != nil {
			logger.Error(errors.ZapUserFetchError, scanError)
			return nil, response.NewInternalServerError(errors.InternalServerErrorMessage)
		}
		users = append(users, user)
	}

	return users, nil
}
