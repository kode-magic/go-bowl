package demographics

import (
	"fmt"
)

func OnCreateCountry() {
	createCountry := Country{
		Name:      "Liberia",
		Continent: "Africa",
	}
	result, err := AddCountry(createCountry)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func GormCountryCreate() {
	createData := Country{
		Name:      "Niger",
		Continent: "Africa",
	}
	result, err := CreateCountry(&createData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func GormCountryList() {
	result, err := ListCountries()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n\n Country count %i", len(*result))
	fmt.Println(result)
}

func OnFetchAllCountries() {
	result := FetchCountries()
	fmt.Println(result)
}
