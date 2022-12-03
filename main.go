package main

import (
	"fmt"
	"log"

	"github.com/saurabh0719/pswHash"
)

func main() {
	// Get the stored password from the database
	storedPassword := "pbkdf2_sha256$260000$1F88YYv36fgyfuglTLvDqA$MGarok+V56arGG70VqvFv4xuCiN5qZCHFu/tkcNAMI8="

	passwordToVerify := "gary123"
	isCorrect := pswHash.Verify(passwordToVerify, storedPassword)
	if !isCorrect {
		log.Fatal("Incorrect password")
	}
	fmt.Println("Correct password")
}
