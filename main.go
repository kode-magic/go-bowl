package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/base64"
	"github.com/kode-magic/go-bowl/validator"
)

func main() {
	fmt.Println("*******************************************************************************")
	fmt.Println("\t\t\t\tGO BOWL")
	fmt.Println("phone: ", validator.IsPhoneNumber("23288110178"))
	fmt.Println("email: ", validator.IsEmail("james@medram.dev"))
	fmt.Println("encoded: ", base64.EncodeString("james@medram.dev"))
	fmt.Println("decoded", base64.DecodeString("amFtZXNAbWVkcmFtLmRldg=="))
}
