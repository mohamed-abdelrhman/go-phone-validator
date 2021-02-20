package countries

import "encoding/json"



type PrivateCountry struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Regex       string `json:"regex"`
}

func (countries Countries)Marshall() []interface{} {
	result :=make([]interface{},len(countries))
	for index,country :=range countries  {
		result[index]=country.Marshall()
	}
	return  result
}

func (country *Country)Marshall() interface{} {
	countryJSON,_:=json.Marshal(country)
	var privateCountry PrivateCountry
	_=json.Unmarshal(countryJSON,&privateCountry)
	return privateCountry
}