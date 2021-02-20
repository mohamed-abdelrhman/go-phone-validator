package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/mohamed-abdelrhman/go-phone-validator/app/domain/customers"
	"github.com/mohamed-abdelrhman/go-phone-validator/services"
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var getCustomerFunc func(customerID int64) (*customers.Customer,*errors.RestErr)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type customersServiceMock struct{}

func (m *customersServiceMock) GetCustomer(i int64) (*customers.Customer, *errors.RestErr) {
	return getCustomerFunc(i)
}

func (m *customersServiceMock) GetAllCustomers(i int64) (customers.Customers, *errors.RestErr) {
	panic("implement me")
}

func (m *customersServiceMock) FilterCustomers(customer customers.FilterCustomer, i int64) (customers.Customers, *errors.RestErr) {
	panic("implement me")
}

func (m *customersServiceMock) CreateCustomer(customer customers.Customer) (*customers.Customer, *errors.RestErr) {
	panic("implement me")
}

func (m *customersServiceMock) UpdateCustomer(b bool, customer customers.Customer) (*customers.Customer, *errors.RestErr) {
	panic("implement me")
}

func (m *customersServiceMock) DeleteCustomer(customerID int64) *errors.RestErr {
	panic("implement me")
}


func TestGetCustomerNotFound(t *testing.T) {
	getCustomerFunc = func(customerID int64) (*customers.Customer,*errors.RestErr) {
		return nil, &errors.RestErr{Status: http.StatusNotFound, Message: "Customer not found"}
	}

	services.CustomerServices = &customersServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/customers/1", nil)
	Get(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

}

