package demographics

import (
	"fmt"
)

func OnCreateCountry() {
	createCountry := Country{
		Nmae: "Liberia",
		Continent: "Africa",
	}
	result, err := AddCountry(createCountry)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func OnFetchAllCountries() {
	result := FetchCountries()
	fmt.Println(result)
}