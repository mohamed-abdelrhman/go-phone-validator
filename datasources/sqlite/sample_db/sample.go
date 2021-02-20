package sample_db

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var(
	SqliteClient *sql.DB
)


func init()  {
	if _, err := os.Stat("./datasources/sqlite/sample_db/sample.db"); os.IsNotExist(err) {
		file, err := os.Create("./datasources/sqlite/sample_db/sample.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}
	SqliteClient, _ = sql.Open("sqlite3", "./datasources/sqlite/sample_db/sample.db") // Open the created SQLite File
	createCountriesTable(SqliteClient) // Create Database Tables
	createCustomersTable(SqliteClient) // Create Database Tables
}

func createCustomersTable(db *sql.DB) {
	//create customer table
	createCustomersTableSQL := `CREATE TABLE customers (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"phone" TEXT,
		"country_id" INTEGER ,
		"valid" TEXT 
	  );`
	statement, err := db.Prepare(createCustomersTableSQL)
	if err != nil {
		log.Println(err.Error())
		return
	}
	statement.Exec()
	log.Println("customers table created")

	//populate customer table
	type customer struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
		CountryID int `json:"country_id"`
		Valid bool `json:"valid"`
	}
	plan, readErr:= ioutil.ReadFile("./datasources/sqlite/sample_db/customer.json")
	if readErr != nil {
		log.Println(readErr)
	}
	var customers []customer
	err = json.Unmarshal(plan, &customers)
	if err != nil {
		log.Println("Cannot unmarshal the json ", err)
	}
	for _,customer :=range customers  {
		insertCustomer(db,customer.Name, customer.Phone,customer.CountryID,customer.Valid)
	}
	log.Println("customers populated !")

}


func createCountriesTable(db *sql.DB) {

	createCountriesTableSQL := `CREATE TABLE countries (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"country_code" TEXT,
		"regex" TEXT		
	  );` // SQL Statement for Create Table
	statement, err := db.Prepare(createCountriesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
		return
	}
		statement.Exec() // Execute SQL Statements
		log.Println("countries table created")
	//populate customer table
	type country struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		CountryCode string `json:"country_code"`
		Regex       string `json:"regex"`
	}
	plan, readErr:= ioutil.ReadFile("./datasources/sqlite/sample_db/countries.json")
	if readErr != nil {
		log.Println(readErr)
	}
	var countries []country
	err = json.Unmarshal(plan, &countries)
	if err != nil {
		log.Println("Cannot unmarshal the json ", err)
	}
	for _,country :=range countries  {
		insertCountry(db,country.Name, country.CountryCode, country.Regex)
	}
	log.Println("countries populated !")


}

// We are passing db reference connection from main to our method with other parameters
func insertCountry(db *sql.DB,  name string, countryCode string, regex string) {
	log.Println("Inserting country record ...")
	insertStudentSQL := `INSERT INTO countries(name, country_code, regex) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, countryCode, regex)
	if err != nil {
		log.Fatalln(err.Error())
	}
}


// We are passing db reference connection from main to our method with other parameters
func insertCustomer(db *sql.DB,  name string, phone string, countryID int, valid bool ) {
	insertStudentSQL := `INSERT INTO customers(name, phone,country_id,valid) VALUES (?, ?,?,?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, phone, countryID,valid)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Populating customer")
}

func getCountry(db *sql.DB, countryCode string)  {
	getCountrySQL := `SELECT * FROM countries WHERE country_code=?`

	stmt,err:= db.Prepare(getCountrySQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()
	var country interface{}
	result:=stmt.QueryRow(countryCode)
	if err:= result.Scan(&country);err!=nil {
		log.Fatal(err.Error())
	}
	log.Println(country)
}