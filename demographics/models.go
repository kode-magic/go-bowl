package demographics

type Country struct {
	Id        string `json:"id"`
	Nmae      string `json:"name"`
	Continent string `json:"continent"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Region struct {
	Id        string  `json:"id"`
	Nmae      string  `json:"name"`
	CountryId string  `json:"countryId"`
	Country   Country `json:"country"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type District struct {
	Id        string  `json:"id"`
	Nmae      string  `json:"name"`
	RegionId  string  `json:"regionId"`
	Region    Region  `json:"region"`
	CountryId string  `json:"countryId"`
	Country   Country `json:"country"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
