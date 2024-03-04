package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// make me a const of the alphabet lowercase and uppercase as a byte array
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// make me a const of the numbers 0-9
const numbers = "0123456789"

// make me a const of the special characters
const special = "!@#$%^&*()_+/'\\{}[]|;:,.<>?"

func generate(length int) {
	password := ""

	charset := alphabet + numbers + special

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password += string(charset[n.Int64()])
	}

	println(password)
}

// get user input for password length
func main() {
	var length int
	print("Enter the length of the password: ")
	fmt.Scan(&length)
	generate(length)
}
