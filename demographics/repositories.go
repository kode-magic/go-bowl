package demographics

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ListLocalDemographics() Demographics {
	content, err := ioutil.ReadFile("./demographics/demo.json")
	if err != nil {
		log.Fatal(err)
	}

	var demographics Demographics
	err = json.Unmarshal(content, &demographics)
	if err != nil {
		log.Fatal(err)
	}

	return demographics
}

func RegionData() []string {
	demographics := ListLocalDemographics()
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

	return list
}

func DistrictsByRegion(region string) []string {
	demographics := ListLocalDemographics()
	var districtData []string

	for _, district := range demographics.Demographics {
		if region == district.Region {
			districtData = append(districtData, district.District)
		}

	}

	return districtData
}
