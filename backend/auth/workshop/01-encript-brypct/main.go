package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedStr, err := encryptToBcrypt("hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed String = %s", hashedStr)
}

// encrypt string kedalam bcrypt
func encryptToBcrypt(str string) (string, error) {
	// Task: Hashing the password with the default cost of 10
	passwordByte := []byte(str)
	hashByte, err := bcrypt.GenerateFromPassword(passwordByte, 10)
	hashStr := string(hashByte)
	return hashStr, err
}
