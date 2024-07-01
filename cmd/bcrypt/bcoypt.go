package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "campare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])

	}

}

func hash(s string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", string(hashedBytes))
}

func compare(s1, s2 string) {
	err := bcrypt.CompareHashAndPassword([]byte(s2), []byte(s1))
	if err != nil {
		fmt.Printf("passwod is invalid %v/n", s1)
		return
	}
	fmt.Printf("passwod is invalid")
}
