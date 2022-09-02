package demographics

func CreateCountry(country *Country) (*Country, error) {
	db, err := ConnectGormDB()
	CheckErr(err)
	createErr := db.Create(&country).Error
	if createErr != nil {
		return nil, createErr
	}

	return country, nil
}

func ListCountries() (*[]Country, error) {
	db, err := ConnectGormDB()
	CheckErr(err)
	var countries []Country
	fetchErr := db.Find(&countries).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return &countries, nil
}

func AddRegion(region *Region) (*Region, error) {
	db, err := ConnectGormDB()
	CheckErr(err)
	createErr := db.Create(&region).Error
	if createErr != nil {
		return nil, createErr
	}

	return region, nil
}

func ListCountryRegions(country string) (*[]Region, error) {
	db, err := ConnectGormDB()
	CheckErr(err)
	var regions []Region
	fetchErr := db.Preload("Country").Where("country_id = ?", country).Find(&regions).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return &regions, nil
}

func GetRegionById(id string) (*Region, error) {
	db, err := ConnectGormDB()
	CheckErr(err)
	var region Region
	
	fetchErr := db.Preload("Country").Where("id = ?", id).Take(&region).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return &region, nil
}