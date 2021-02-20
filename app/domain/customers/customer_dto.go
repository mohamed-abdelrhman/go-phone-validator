package customers

import (
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/errors"
	"strings"
)

const (
	StatusActive="active"
)

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CountryID int64 `json:"country_id"`
	Valid    bool `json:"valid"`
}

type FilterCustomer struct {
	CountryID int64 `json:"country_id"`
	Valid    string `json:"valid"`
}
type Customers []Customer

func ( customer *Customer) Validate() *errors.RestErr {
	customer.Name=strings.TrimSpace(strings.ToLower(customer.Name))
	customer.Phone=strings.TrimSpace(strings.ToLower(customer.Phone))
	//validate Inputs

	return  nil
}