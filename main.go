package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/demographics"
)

func main() {
	// getGst := calculations.CalculateGST(9000, 15)
	fmt.Println("Hello from GO-BOWL")

	dataDistrict := demographics.DistrictsByRegion("Northern")
	dataRegion := demographics.RegionData()
	demographics.ListLocalDemographics()
	fmt.Println(dataRegion)
	fmt.Println(dataDistrict)
}
