package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mohamed-abdelrhman/phoneValidator/utils/errors"
	"strings"
)

const (
	errorNoRows="no rows in result set"
) 

func ParseError(err error)*errors.RestErr  {
	sqlErr,ok:=err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(),errorNoRows) {
			return  errors.NewNotFoundError("No Such Record")
		}
		return  errors.NewInternalServerError("Error Parsing Database Response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data duplicated key")
	}
	return  errors.NewInternalServerError("error processing request")
}