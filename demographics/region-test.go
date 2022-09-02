package demographics

import "fmt"

func OnAddRegion() {
	createRegion := Region{
		Nmae: "South",
		CountryId: "42c441a0-2ad5-11ed-ac47-acde48001122",
	}
	result, err := AddCountryRegion(createRegion)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func OnFetchCountryRegions(country string) {
	result := FetchCountryRegions(country)
	fmt.Println(result)
}