package demographics

func CreateCountry(country *Country) (*Country, error) {
	db, err := connectGormDB()
	CheckErr(err)
	createErr := db.Create(&country).Error
	if createErr != nil {
		return nil, createErr
	}

	return country, nil
}

func ListCountries() (*[]serialCountry, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var countries []Country
	fetchErr := db.Find(&countries).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	serialCountries := make([]serialCountry, len(countries))
	for i, country := range countries {
		serialCountries[i] = *country.serializeCountry()
	}

	return &serialCountries, nil
}

func AddRegion(region *Region) (*Region, error) {
	db, err := connectGormDB()
	CheckErr(err)
	createErr := db.Create(&region).Error
	if createErr != nil {
		return nil, createErr
	}

	return region, nil
}

func ListCountryRegions(country string) (*[]serialRegion, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var regions []Region
	fetchErr := db.Preload("Country").Where("country_id = ?", country).Find(&regions).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	serialRegions := make([]serialRegion, len(regions))
	for i, region := range regions {
		serialRegions[i] = *region.serializeRegion()
	}

	return &serialRegions, nil
}

func GetRegionById(id string) (*serialRegion, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var region Region

	fetchErr := db.Preload("Country").Where("id = ?", id).Take(&region).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return region.serializeRegion(), nil
}

func GetRegionByName(name string) (*serialRegion, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var region Region

	fetchErr := db.Preload("Country").Where("name = ?", name).Take(&region).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return region.serializeRegion(), nil
}

func AddDistrict(district *District) (*District, error) {
	db, err := connectGormDB()
	CheckErr(err)
	createErr := db.Create(&district).Error
	if createErr != nil {
		return nil, createErr
	}

	return district, nil
}

func GetDistrictById(id string) (*serialDistrict, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var district District

	fetchErr := db.Preload("Region").Where("id = ?", id).Take(&district).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	return district.serializeDistrict(), nil
}

func ListRegionDistricts(region string) (*[]serialDistrict, error) {
	db, err := connectGormDB()
	CheckErr(err)
	var districts []District
	fetchErr := db.Preload("Region").Where("region_id = ?", region).Find(&districts).Error
	if fetchErr != nil {
		return nil, fetchErr
	}

	serialDistricts := make([]serialDistrict, len(districts))
	for i, district := range districts {
		serialDistricts[i] = *district.serializeDistrict()
	}


	return &serialDistricts, nil
}
