package simpro

// CompanyListResponse
type CompanyListResponse struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}

// CompanyResponse
type CompanyResponse struct {
	ID       uint           `json:"ID"`
	Name     string         `json:"Name"`
	Phone    string         `json:"Phone"`
	Email    string         `json:"Email"`
	Address  companyAddress `json:"Address"`
	Country  string         `json:"Country"`
	Timezone string         `json:"Timezone"`
	Currency string         `json:"Currency"`
}

// companyAddress
type companyAddress struct {
	Line1 string `json:"Line1"`
	Line2 string `json:"Line2"`
}
