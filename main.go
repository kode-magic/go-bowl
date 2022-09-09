package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/demographics"
)

func main() {
	fmt.Println("*******************************************************************************")
	fmt.Println("\t\t\t\tGO BOWL")
	demographics.DistrictsByRegion("Western")
	fmt.Println("*******************************************************************************")
}
