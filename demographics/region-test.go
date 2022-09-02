package demographics

import "fmt"

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
