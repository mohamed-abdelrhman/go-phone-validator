package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/phoneValidator/domain/countries"
	"github.com/mohamed-abdelrhman/phoneValidator/services"
	"github.com/mohamed-abdelrhman/phoneValidator/utils/errors"
	"github.com/mohamed-abdelrhman/phoneValidator/utils/logger"
	"github.com/mohamed-abdelrhman/phoneValidator/utils/strings_utils"
	"net/http"
	"strconv"
)

var(
	CountryControllers countryCountryInterface =&countryControllers{}
)

type countryControllers struct {}
type countryCountryInterface interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func  (co *countryControllers)Create(c *gin.Context)  {
	var country countries.Country
	if err:=c.ShouldBindJSON(&country);err!=nil{
		restErr:=errors.NewBadRequestError("Invalid Form Body")
		logger.Error("error Binding save request",err)
		c.JSON(restErr.Status,restErr)
		return
	}
	result,saveError := services.CountryServices.CreateCountry(country)
	if saveError != nil {
		c.JSON(saveError.Status,saveError)
		return
	}
	c.JSON(http.StatusCreated,result.Marshall())
}

func (co *countryControllers)Get(c *gin.Context)  {
	countryID,idErr:= strings_utils.ParseId(c.Param("country_id"),"country Id")

	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
		return
	}
	country,getErr:= services.CountryServices.GetCountry(countryID)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK,country.Marshall())

}

func (co *countryControllers)GetAll(c *gin.Context)  {
	pageNo,pageErr:=strconv.ParseInt(c.Query("page"),10,64)
	if pageErr !=nil{
		pageNo=0
	}
	AllCountries,err:= services.CountryServices.GetAllCountries(pageNo)
	if err != nil {
		c.JSON(err.Status,err)
	}
	c.JSON(http.StatusOK, AllCountries.Marshall())


}
func (co *countryControllers)Update(c *gin.Context)  {
	countryID,idErr:= strings_utils.ParseId(c.Param("country_id"),"country Id")
	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
		return
	}
	var country countries.Country
	if  err:=c.ShouldBindJSON(&country);err !=nil{
		restErr:=errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	country.ID =countryID

	isPartial:=c.Request.Method==http.MethodPatch

	result,err:= services.CountryServices.UpdateCountry(isPartial,country)
	if err!=nil {
		c.JSON(err.Status,err)
	}
	c.JSON(http.StatusOK,result.Marshall())
}

func (co *countryControllers)Delete(c *gin.Context)  {
	countryID,idErr:= strings_utils.ParseId(c.Param("country_id"),"country Id")
	if idErr!=nil {
		c.JSON(idErr.Status,idErr)
	}
	if err:= services.CountryServices.DeleteCountry(countryID);err!=nil {
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,map[string]string{"status":"deleted"})
}

