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

