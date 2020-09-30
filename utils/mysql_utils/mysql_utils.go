package mysql_utils

import (
	"fmt"
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
			return errors.NewNotFoundError(fmt.Sprintf(errorUserIdNotFound))
		}
		return errors.NewInternalServerError(fmt.Sprintf(errorGenericMessage, err.Error()))
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf(errorDuplicateData, sqlError.Message))
	}
	return errors.NewInternalServerError(
		fmt.Sprintf(errorGenericMySQL, sqlError.Number, sqlError.Message, sqlError.Error()))
}
