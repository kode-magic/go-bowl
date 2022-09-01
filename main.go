package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/calculations"
	"github.com/kode-magic/go-bowl/ulid"
)

func main () {
	getGst := calculations.CalculateGST(9000, 15)
	fmt.Println(getGst)

	noGst := calculations.RemoveGST(getGst.AmountWithGst, 15)
	fmt.Println(noGst)

	id := ulid.GenerateUUID()
	fmt.Println("uuid = "+id.String())
}