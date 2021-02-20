package services

import (
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/domain/countries"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/errors"
)

var(
	CountryServices countryServiceInterface =&countryServices{}
)

type countryServices struct {}

type countryServiceInterface interface {
	GetCountry(int64)  (*countries.Country,*errors.RestErr)
	GetAllCountries(int64)  (countries.Countries, *errors.RestErr)
	CreateCountry(countries.Country)(*countries.Country,*errors.RestErr)
	UpdateCountry(bool, countries.Country)(*countries.Country,*errors.RestErr)
	DeleteCountry(countryID int64)*errors.RestErr
}

func (s *countryServices) GetCountry(countryID int64)  (*countries.Country,*errors.RestErr)  {
	result:=&countries.Country{ID: countryID}
	if err:=result.Get();err !=nil {
		return nil,err
	}
	return result,nil
}

func (s *countryServices) GetAllCountries(pageNo int64)(countries.Countries, *errors.RestErr)  {
	dao:=&countries.Country{}
	return dao.GetAll(pageNo)
}

func (s *countryServices) CreateCountry(country countries.Country)(*countries.Country,*errors.RestErr)  {
	//add another country fields like date created
	//country.DateCreated=date.GetNowDBFormat()
	if err:=country.Save();err!=nil{
		return nil,err
	}
	return &country,nil
}

func (s *countryServices) UpdateCountry(isPartial bool,country countries.Country)(*countries.Country,*errors.RestErr) {
	current,err:= CountryServices.GetCountry(country.ID)
	if err!=nil{
		return  nil,err
	}
	if isPartial {
		if country.Name !=""{
			current.Name=country.Name
		}
		if country.CountryCode !=""{
			current.CountryCode=country.CountryCode
		}
		if country.Regex !=""{
			current.Regex=country.Regex
		}
	}else {
		current.Name=country.Name
		current.CountryCode=country.CountryCode
		current.Regex=country.Regex
	}

	if err:=current.Update();err!=nil{
		return nil,err
	}
	return current,nil



}

func (s *countryServices) DeleteCountry(countryID int64)*errors.RestErr  {
	country:=&countries.Country{ID: countryID}
	return country.Delete()
}