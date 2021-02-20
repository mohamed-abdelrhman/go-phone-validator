package countries

type Country struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Regex       string `json:"regex"`
}
type Countries []Country