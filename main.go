package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

// generatePassword generates a random password with the specified length and complexity
func generatePassword(length int, uppercase bool, lowercase bool, digits bool, specialChars bool) (string, error) {
	var chars string

	if uppercase {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if lowercase {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}

	if digits {
		chars += "0123456789"
	}

	if specialChars {
		chars += "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
	}

	if len(chars) == 0 {
		return "", fmt.Errorf("at least one character set (uppercase, lowercase, digits, specialChars) must be selected")
	}

	var password string
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password += string(chars[randomIndex.Int64()])
	}

	return password, nil
}

func main() {
	// Command-line flags
	length := flag.Int("length", 12, "Length of the password")
	uppercase := flag.Bool("uppercase", true, "Include uppercase letters")
	lowercase := flag.Bool("lowercase", true, "Include lowercase letters")
	digits := flag.Bool("digits", true, "Include digits")
	specialChars := flag.Bool("specialChars", true, "Include special characters")
	flag.Parse()

	// Generate the password
	password, err := generatePassword(*length, *uppercase, *lowercase, *digits, *specialChars)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the generated password
	fmt.Println("Generated Password:", password)
}

// go run main.go -length 16 -uppercase -lowercase -digits -specialChars
