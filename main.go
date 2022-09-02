package main

import (
	"github.com/kode-magic/go-bowl/demographics"
)

func main () {
	// getGst := calculations.CalculateGST(9000, 15)
	// fmt.Println(getGst)

	// noGst := calculations.RemoveGST(getGst.AmountWithGst, 15)
	// fmt.Println(noGst)

	// id := ulids.GenerateUUID()
	// fmt.Println("uuid = "+id.String())

	// demographics.OnFetchCountryRegions("42c441a0-2ad5-11ed-ac47-acde48001122")
	// demographics.GormCountryCreate()
	// demographics.GormCountryList()
	// demographics.GormAddRegion()
	// demographics.GormCountryRegions()
	demographics.GormRegionById()

}