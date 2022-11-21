package main

import (
	"fmt"
	"github.com/kode-magic/go-bowl/bcrypt"
	"github.com/kode-magic/go-bowl/jwt"
	"time"
)

func main() {
	fmt.Println("*******************************************************************************")
	fmt.Println("\t\t\t\tGO BOWL")
	hash, _ := bcrypt.HashPassword("Kodemagik!!1")
	fmt.Println("hashes: ", hash)
	fmt.Println("compare: ", bcrypt.CheckPasswordHash("Kodemagik!!1", hash))
	fmt.Println("token: ", jwt.GenerateJWTForEmail("james@medram.dev", time.Hour*24))
}
