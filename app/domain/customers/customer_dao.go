package customers

import (
	"database/sql"
	"github.com/mohamed-abdelrhman/go-phone-validator/datasources/sqlite/sample_db"
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/errors"
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/logger"
	"strconv"
)

const(
	queryInsertCustomer="INSERT INTO customers(name,phone,country_id,valid) VALUES(?,?,?,?);"
	queryGetCustomer=" SELECT * FROM customers WHERE id=?;"
	queryGetAllCustomer=" SELECT * from customers LIMIT ?,? ;"
	queryGetCustomersByCountryIdAndStatus=" SELECT * from customers WHERE country_id=? AND valid=? LIMIT ?,? ;"
	queryGetCustomersByCountryId=" SELECT * from customers WHERE country_id=? LIMIT ?,? ;"
	queryGetCustomersByStatus=" SELECT * from customers WHERE valid=? LIMIT ?,? ;"
	queryUpdateCustomer="UPDATE customers SET name=?,phone=?,country_id=?,valid=? WHERE id=?;"
	queryDeleteCustomer="DELETE FROM customers WHERE id=?;"
	)

var pageLimit int64=10


func (customer *Customer)Get() *errors.RestErr {
	stmt,err:= sample_db.SqliteClient.Prepare(queryGetCustomer)
	if err != nil {
		logger.Error("error where trying get prepare customer stmt",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	result:=stmt.QueryRow(customer.ID)
	if err:= result.Scan(&customer.ID,&customer.Name,&customer.Phone,&customer.CountryID,&customer.Valid);err!=nil {
		logger.Error("error where trying get query row customer ",err)
		return  errors.NewInternalServerError("Database parsing error")	}
	return  nil
}

func (customer *Customer)GetAll(pageNo int64)([]Customer, *errors.RestErr){
	stmt, err :=sample_db.SqliteClient.Prepare(queryGetAllCustomer)
	if err != nil {
		logger.Error("error where trying All find customers ",err)
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	rows,err:=stmt.Query(pageLimit*(pageNo-1), pageLimit)
	if err != nil {
		return nil, errors.NewInternalServerError("Database parsing error")

	}
	defer rows.Close()
	results:=make([]Customer,0)
	for rows.Next(){
		var customer Customer
		if err :=rows.Scan(&customer.ID,&customer.Name,&customer.Phone, &customer.CountryID, &customer.Valid);err!=nil{
			logger.Error("error where trying to find All Customers ",err)
			return nil, errors.NewInternalServerError("Database parsing error")
		}
		results=append(results,customer)
	}
	if len(results)==0 {
		return  nil, errors.NewNotFoundError("no Customers Found")
	}
	return results,nil



}
func (customer *Customer)Filter(filterCustomer FilterCustomer, pageNo int64)([]Customer, *errors.RestErr){

	valid,parsValidErr:=strconv.ParseInt(filterCustomer.Valid,10,64)

	var stmt *sql.Stmt
	var err error
	if filterCustomer.CountryID != 0 && parsValidErr == nil {
		if valid==1 {customer.Valid =true}else{ customer.Valid= false}
		stmt, err =sample_db.SqliteClient.Prepare(queryGetCustomersByCountryIdAndStatus)
	}else if filterCustomer.CountryID == 0 && parsValidErr != nil {
		stmt, err =sample_db.SqliteClient.Prepare(queryGetAllCustomer)
	}else if  filterCustomer.CountryID == 0 && parsValidErr == nil {
		if valid==1 {customer.Valid =true}else{ customer.Valid= false}
		stmt, err =sample_db.SqliteClient.Prepare(queryGetCustomersByStatus)
	}else if  filterCustomer.CountryID != 0 && parsValidErr != nil {
		stmt, err =sample_db.SqliteClient.Prepare(queryGetCustomersByCountryId)
	}

	if err != nil {
		logger.Error("error where trying All find customers ",err)
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()

	var rows *sql.Rows

	if filterCustomer.CountryID != 0 && parsValidErr == nil {
		rows,err=stmt.Query(filterCustomer.CountryID,customer.Valid, pageLimit*(pageNo-1), pageLimit)
	}else if filterCustomer.CountryID == 0 && parsValidErr != nil {
		rows,err=stmt.Query(pageLimit*(pageNo-1), pageLimit)
	}else if  filterCustomer.CountryID == 0 && parsValidErr == nil {
		rows,err=stmt.Query(customer.Valid, pageLimit*(pageNo-1), pageLimit)
	}else if  filterCustomer.CountryID != 0 && parsValidErr != nil {
		rows,err=stmt.Query(filterCustomer.CountryID, pageLimit*(pageNo-1), pageLimit)
	}



	if err != nil {
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer rows.Close()
	results:=make([]Customer,0)
	for rows.Next(){
		var customer Customer
		if err :=rows.Scan(&customer.ID,&customer.Name,&customer.Phone, &customer.CountryID, &customer.Valid);err!=nil{
			logger.Error("error where trying to find All Customers ",err)
			return nil, errors.NewInternalServerError("Database parsing error")
		}
		results=append(results,customer)
	}
	if len(results)==0 {
		return  nil, errors.NewNotFoundError("no Customers Found")
	}
	return results,nil



}

func (customer *Customer)Save() *errors.RestErr {
	stmt,err := sample_db.SqliteClient.Prepare(queryInsertCustomer)
	if err != nil {
		logger.Error("error where trying get prepare customer stmt",err)
		return errors.NewInternalServerError("database error")
	}
	defer  stmt.Close()
	insertResult,err:=stmt.Exec(customer.Name,customer.Phone,customer.CountryID,customer.Valid)
	if err !=nil {
		logger.Error("error where trying get save customer ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	customerID,err:=insertResult.LastInsertId()
	if err != nil {
		logger.Error("error where trying get last customer Id ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	customer.ID = customerID
	return nil
}

func (customer *Customer)Update()*errors.RestErr {

	stmt,err:= sample_db.SqliteClient.Prepare(queryUpdateCustomer)
	if err != nil {
		logger.Error("error where trying prepare customer update ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	_,err=stmt.Exec(customer.Name,customer.Phone,customer.CountryID,customer.Valid,customer.ID)

	if err != nil {
		logger.Error("error where trying execute customer update",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	return nil
}

func (customer *Customer)Delete()*errors.RestErr {

	stmt,err:= sample_db.SqliteClient.Prepare(queryDeleteCustomer)
	if err != nil {
		logger.Error("error where trying prepare customer delete ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	_,err=stmt.Exec(customer.ID)
	if err != nil {
		logger.Error("error where trying get execute customer delete ",err)
		return  errors.NewInternalServerError("Database parsing error")
	}
	return nil
}
