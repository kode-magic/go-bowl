package demographics

import (
	"encoding/json"
	"fmt"
	"log"
)

func LocalDemographics() Demographics {
	var demographics Demographics
	err := json.Unmarshal([]byte(DemographicsJSON), &demographics)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(demographics)

	return demographics
}

func RegionData() []string {
	demographics := LocalDemographics()
	var regions []string

	for _, district := range demographics.Demographics {
		regions = append(regions, district.Region)
	}

	allKeys := make(map[string]bool)
	var list []string
	for _, item := range regions {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	fmt.Println(list)

	return list
}

func DistrictsByRegion(region string) []string {
	demographics := LocalDemographics()
	var districtData []string

	for _, district := range demographics.Demographics {
		if region == district.Region {
			districtData = append(districtData, district.District)
		}

	}

	return districtData
}
