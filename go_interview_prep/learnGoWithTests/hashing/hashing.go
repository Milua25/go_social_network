package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	hashedPassword := sha256.Sum256([]byte(password))
	fmt.Printf("Hashed Password: %x\n", hashedPassword)

	hashedPassword = [32]byte(sha256.New().Sum([]byte(password)))
	fmt.Printf("Hashed Password 2: %x\n", hashedPassword)

	// salt the password

	salt, err := generateSalt()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Salt: %x\n", salt)

	newHashedPassword, err := hashPassword(password, salt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New Hashed Password: %s\n", newHashedPassword)

	// store salt and hashed password in a database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("Salt String: %s\n", saltStr)
	fmt.Printf("Hashed Password: %s\n", newHashedPassword)

	// retrieve salt and hashed password from the database
	retrievedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved Salt: %x\n", retrievedSalt)

	retrievedHashedPassword, err := hashPassword(password, retrievedSalt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved Hashed Password: %s\n", retrievedHashedPassword)

	if retrievedHashedPassword == newHashedPassword {
		fmt.Println("Passwords match")
	} else {
		fmt.Println("Passwords do not match")
	}
}

func generateSalt() ([]byte, error) {

	saltSlice := make([]byte, 16)

	_, err := io.ReadFull(rand.Reader, saltSlice)
	if err != nil {
		return nil, err
	}

	return saltSlice, nil
}

func hashPassword(password string, salt []byte) (string, error) {

	hash := sha256.New().Sum([]byte(password + string(salt)))
	return base64.StdEncoding.EncodeToString(hash), nil
}
