package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/calculations"
)

func main () {
	getGst := calculations.CalculateGST(9000, 15)
	fmt.Printf("Include GST = %f", getGst)

	noGst := calculations.RemoveGST(getGst, 15)
	fmt.Printf("Remove GST = %f", noGst)
}