package mysql_utils

import (
	"fmt"
	"github.com/adershrp/bookstore_users-api/logger"
	"github.com/adershrp/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	noRowsInResult = "no rows in result set"
)

const (
	errorGenericMessage = "Database Error : %s"
	errorDuplicateData  = "Invalid data: %s"
	errorGenericMySQL   = "Database Error : %d, %s, %s"
	errorUserIdNotFound = "No record matching given Id."
)

/**
Generic Error Parser
*/
func ParseError(err error) *errors.RestError {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsInResult) {
			logger.Error("No result found", err)
			return errors.NewNotFoundError(fmt.Sprintf(errorUserIdNotFound))
		}
		logger.Error("Internal Server Error.", err)
		return errors.NewInternalServerError(fmt.Sprintf(errorGenericMessage, err.Error()))
	}
	switch sqlError.Number {
	case 1062:
		logger.Error("Bad request payload, duplicate record.", sqlError)
		return errors.NewBadRequestError(fmt.Sprintf(errorDuplicateData, sqlError.Message))
	}
	logger.Error("Internal Server Error.", sqlError)
	return errors.NewInternalServerError(
		fmt.Sprintf(errorGenericMySQL, sqlError.Number, sqlError.Message, sqlError.Error()))
}
