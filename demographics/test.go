package demographics

import "fmt"

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

func GormAddRegion() {
	createRegion := Region{
		Name:      "West",
		CountryId: "42c441a0-2ad5-11ed-ac47-acde48001122",
	}
	result, err := AddRegion(&createRegion)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func GormCountryRegions() {
	result, err := ListCountryRegions("42c441a0-2ad5-11ed-ac47-acde48001122")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func GormRegionById() {
	result, err := GetRegionById("1d9529f6-2b09-11ed-85d3-acde48001122")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func GormRegionByName() {
	result, err := GetRegionByName("West")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func GormAddDistrict() {
	createDistrict := District{
		Name:      "Bombali",
		RegionId: "c8be22c4-2aeb-11ed-b3f5-acde48001122",
	}
	result, err := AddDistrict(&createDistrict)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func GormRegionDistricts() {
	result, err := ListRegionDistricts("c8be22c4-2aeb-11ed-b3f5-acde48001122")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
