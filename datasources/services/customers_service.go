package services

import (
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/domain/countries"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/domain/customers"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/errors"
	"regexp"
)

var(
	CustomerServices customerServiceInterface =&customerServices{}
)

type customerServices struct {}

type customerServiceInterface interface {
	GetCustomer(int64)  (*customers.Customer,*errors.RestErr)
	GetAllCustomers(int64)  (customers.Customers, *errors.RestErr)
	FilterCustomers(customers.FilterCustomer,int64)  (customers.Customers, *errors.RestErr)
	CreateCustomer(customers.Customer)(*customers.Customer,*errors.RestErr)
	UpdateCustomer(bool, customers.Customer)(*customers.Customer,*errors.RestErr)
	DeleteCustomer(customerID int64)*errors.RestErr
}

func (s *customerServices) GetCustomer(customerID int64)  (*customers.Customer,*errors.RestErr)  {
	result:=&customers.Customer{ID: customerID}
	if err:=result.Get();err !=nil {
		return nil,err
	}
	return result,nil
}

func (s *customerServices) GetAllCustomers(pageNo int64)(customers.Customers, *errors.RestErr)  {
	dao:=&customers.Customer{}
	return dao.GetAll(pageNo)
}


func (s *customerServices) FilterCustomers(customer customers.FilterCustomer,pageNo int64)(customers.Customers, *errors.RestErr)  {

	dao:=customers.Customer{}
	return dao.Filter(customer,pageNo)
}

func (s *customerServices) CreateCustomer(customer customers.Customer)(*customers.Customer,*errors.RestErr)  {
	if err :=customer.Validate();err!=nil {
		return   nil,err
	}
	// get country
	country:=&countries.Country{ID: customer.CountryID}
	if err:=country.Get();err!=nil{
		return nil,err
	}
	// validate phone regex
	reg := regexp.MustCompile(country.Regex)
	customer.Valid = false
	if reg.MatchString(customer.Phone) {
		customer.Valid = true
	}
	if err:=customer.Save();err!=nil{
		return nil,err
	}
	return &customer,nil
}

func (s *customerServices) UpdateCustomer(isPartial bool,customer customers.Customer)(*customers.Customer,*errors.RestErr) {
	current,err:= CustomerServices.GetCustomer(customer.ID)
	if err!=nil{
		return  nil,err
	}
	if isPartial {
		if customer.Name !=""{
			current.Name=customer.Name
		}
		if customer.Phone !=""{
			current.Phone=customer.Phone
		}
		if customer.CountryID !=0 {
			current.CountryID=customer.CountryID
		}
	}else {
		current.Name=customer.Name
		current.Phone=customer.Phone
		current.CountryID=customer.CountryID
	}

	// get country
	country:=&countries.Country{ID: customer.CountryID}
	if err:=country.Get();err!=nil{
		return nil,err
	}
	// validate phone regex
	reg := regexp.MustCompile(country.Regex)
	current.Valid = false
	if reg.MatchString(current.Phone) {
		current.Valid = true
	}

	if err:=current.Update();err!=nil{
		return nil,err
	}

	return current,nil
}

func (s *customerServices) DeleteCustomer(customerID int64)*errors.RestErr  {
	customer:=&customers.Customer{ID: customerID}
	return customer.Delete()
}