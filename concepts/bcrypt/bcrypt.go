package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password1234`
	loginPwd := `password1234`
	encryptedString, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	fmt.Println(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encryptedString)
	err = bcrypt.CompareHashAndPassword(encryptedString, []byte(loginPwd))
	if err != nil {
		fmt.Println("Can't login...")
	}
}
