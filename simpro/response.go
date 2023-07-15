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

// SecurityGroupListResponse
type SecurityGroupListResponse struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}

// SecurityGroupResponse
type SecurityGroupResponse struct {
	ID            uint                        `json:"ID"`
	Name          string                      `json:"Name"`
	BusinessGroup *securityGroupBusinessGroup `json:"BusinessGroup"`
	Dashboards    []*securityGroupDashboard   `json:"Dashboards"`
}

// securityGroupBusinessGroup
type securityGroupBusinessGroup struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}

// securityGroupDashboard
type securityGroupDashboard struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}
