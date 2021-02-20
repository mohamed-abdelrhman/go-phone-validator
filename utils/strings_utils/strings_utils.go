package strings_utils

import (
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/errors"
	"strconv"
)

func ParseId(Id string,IdName string)(int64,*errors.RestErr)  {
	ParsedID,err:=strconv.ParseInt(Id,10,64)
	if err !=nil{
		return 0, errors.NewBadRequestError("invalid "+IdName)
	}
	return ParsedID,nil
}
