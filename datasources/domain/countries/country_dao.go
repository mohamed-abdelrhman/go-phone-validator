package countries

import (
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/sqlite/sample_db"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/errors"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/utils/logger"
)

const(
	queryInsertCountry="INSERT INTO countries(name,country_code,regex) VALUES(?,?,?);"
	queryGetCountry=" SELECT id,name,country_code,regex FROM countries WHERE id=?"
	queryGetAllCountry=" SELECT * from countries LIMIT ?,? ;"
	queryUpdateCountry="UPDATE countries SET name=?,country_code=?,regex=? WHERE id=?;"
	queryDeleteCountry="DELETE FROM countries WHERE id=?;"
	)

var pageLimit int64=10

func (country *Country)Get() *errors.RestErr {
	stmt,err:= sample_db.SqliteClient.Prepare(queryGetCountry)
	if err != nil {
		logger.Error("error where trying get prepare user stmt",err)
		return  errors.NewInternalServerError("Database parsing error")

	}
	defer stmt.Close()
	result:=stmt.QueryRow(country.ID)
	if err:= result.Scan(&country.ID,&country.Name,&country.CountryCode,&country.Regex);err!=nil {
		logger.Error("error where trying get query row country ",err)
		return  errors.NewInternalServerError("Database parsing error")	}
	return  nil
}
func (country *Country)GetAll(pageNo int64)([]Country, *errors.RestErr){
	stmt, err :=sample_db.SqliteClient.Prepare(queryGetAllCountry)
	if err != nil {
		logger.Error("error where trying All find countries ",err)
		return nil,errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	rows,err:=stmt.Query(pageLimit*(pageNo-1),pageLimit)
	if err != nil {
		return nil,errors.NewInternalServerError("Database parsing error")
	}
	defer rows.Close()
	results:=make([]Country,0)
	for rows.Next(){
		var country Country
		if err :=rows.Scan(&country.ID,&country.Name,&country.CountryCode,&country.Regex);err!=nil{
			logger.Error("error where trying to find All Countries ",err)
			return nil,errors.NewInternalServerError("Database parsing error")
		}
		results=append(results,country)
	}
	if len(results)==0 {
		return  nil,errors.NewNotFoundError("no Countries Found")
	}
	return results,nil



}


func (country *Country)Save() *errors.RestErr  {
	stmt,err := sample_db.SqliteClient.Prepare(queryInsertCountry)
	if err != nil {
		logger.Error("error where trying get prepare country stmt",err)
		return errors.NewInternalServerError("database error")
	}
	defer  stmt.Close()
	insertResult,err:=stmt.Exec(country.Name,country.CountryCode,country.Regex)
	if err !=nil {
		logger.Error("error where trying get save country ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	countryID,err:=insertResult.LastInsertId()
	if err != nil {
		logger.Error("error where trying get last country Id ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	country.ID = countryID
	return nil
}

func (country *Country)Update()*errors.RestErr  {

	stmt,err:= sample_db.SqliteClient.Prepare(queryUpdateCountry)
	if err != nil {
		logger.Error("error where trying prepare country update ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	_,err=stmt.Exec(country.Name,country.CountryCode,country.Regex,country.ID)

	if err != nil {
		logger.Error("error where trying execute country update",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	return nil
}

func (country *Country)Delete()*errors.RestErr  {

	stmt,err:= sample_db.SqliteClient.Prepare(queryDeleteCountry)
	if err != nil {
		logger.Error("error where trying prepare country delete ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	_,err=stmt.Exec(country.ID)
	if err != nil {
		logger.Error("error where trying get execute country delete ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	return nil
}
