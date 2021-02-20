package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/domain/customers"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/services"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/errors"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/logger"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/strings_utils"
	"log"
	"net/http"
	"strconv"
)


var(
	CustomerControllers customerControllerInterface =&customerControllers{}
)

type customerControllers struct {}
type customerControllerInterface interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Filter(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}


func (co *customerControllers)Create(c *gin.Context)  {
	var customer customers.Customer
	if err:=c.ShouldBindJSON(&customer);err!=nil{
		restErr:=errors.NewBadRequestError("Invalid Form Body")
		logger.Error("error Binding save request",err)
		c.JSON(restErr.Status,restErr)
		return
	}
	result,saveError := services.CustomerServices.CreateCustomer(customer)
	if saveError != nil {
		c.JSON(saveError.Status,saveError)
		return
	}
	c.JSON(http.StatusCreated,result.Marshall())
}

func  (co *customerControllers)Get(c *gin.Context)  {
	customerID,idErr:= strings_utils.ParseId(c.Param("customer_id"),"customer Id")
	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
		return
	}
	customer,getErr:= services.CustomerServices.GetCustomer(customerID)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK,customer.Marshall())

}

func  (co *customerControllers)GetAll(c *gin.Context)  {
	pageNo,pageErr:=strconv.ParseInt(c.Query("page"),10,64)
	if pageErr !=nil{
		pageNo=0
	}
	AllCustomers,err:= services.CustomerServices.GetAllCustomers(pageNo)
	if err != nil {
		c.JSON(err.Status,err)
	}
	c.JSON(http.StatusOK, AllCustomers.Marshall())


}
func  (co *customerControllers)Filter(c *gin.Context)  {
	pageNo,pageErr:=strconv.ParseInt(c.Query("page"),10,64)
	if pageErr !=nil{
		pageNo=0
	}
	var customer customers.FilterCustomer
	if err:=c.ShouldBindJSON(&customer);err!=nil{
		restErr:=errors.NewBadRequestError("Invalid Form Body")
		logger.Error("error Binding save request",err)
		c.JSON(restErr.Status,restErr)
		return
	}

	log.Println(customer)

	result,err:= services.CustomerServices.FilterCustomers(customer,pageNo )
	if err != nil {
		c.JSON(err.Status,err)
	}
	c.JSON(http.StatusOK, result.Marshall())

}
func  (co *customerControllers)Update(c *gin.Context)  {
	customerID,idErr:= strings_utils.ParseId(c.Param("customer_id"),"customer Id")
	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
		return
	}
	var customer customers.Customer
	if  err:=c.ShouldBindJSON(&customer);err !=nil{
		restErr:=errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	customer.ID = customerID

	isPartial:=c.Request.Method==http.MethodPatch

	result,err:= services.CustomerServices.UpdateCustomer(isPartial,customer)
	 if err!=nil {
		c.JSON(err.Status,err)
	 }
	 c.JSON(http.StatusOK,result.Marshall())
}

func  (co *customerControllers)Delete(c *gin.Context)  {
	customerID,idErr:= strings_utils.ParseId(c.Param("customer_id"),"customer Id")
	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
	}
	if err:= services.CustomerServices.DeleteCustomer(customerID);err!=nil {
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,map[string]string{"status":"deleted"})
}

