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

