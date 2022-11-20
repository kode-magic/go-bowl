package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/jwt"
	"time"
)

func main() {
	fmt.Println("*******************************************************************************")
	fmt.Println("\t\t\t\tGO BOWL")
	fmt.Println("token: ", jwt.GenerateJWTForEmail("james@medram.dev", time.Hour*24))
}
