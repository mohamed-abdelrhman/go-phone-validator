package customers

import "encoding/json"

type PublicCustomer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CountryID int64 `json:"country_id"`
	Valid    bool `json:"valid"`
}

func (customers Customers)Marshall() []interface{} {
	result :=make([]interface{},len(customers))
	for index,customer :=range customers  {
		result[index]=customer.Marshall()
	}
	return  result
}

func (customer *Customer)Marshall() interface{} {
	customerJSON,_:=json.Marshal(customer)
	var publicCustomer PublicCustomer
	_=json.Unmarshal(customerJSON,&publicCustomer)
	return publicCustomer
}