// main.go
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pratyakshroy47/go-file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	// Check if the command-line argument is provided
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	// Get the command-line argument
	function := os.Args[1]

	// Perform action based on the command-line argument
	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt the file, and decrypt to decrypt the file.")
		os.Exit(1)
	}
}

// Print help information
func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("Simple file encipher for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTrie to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

// Handle file encryption
func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nFile successfully protected...")
}

// Handle file decryption
func decryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter your Password: ")
	password, _ := term.ReadPassword(0)

	fmt.Print("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted...")
}

// Get password input from the user
func getPassword() []byte {
	fmt.Println("Enter your password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again\n")
		return getPassword()
	}

	return password
}

// Validate file existence
func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// Validate password match
func validatePassword(password1 []byte, password2 []byte) bool {
	return bytes.Equal(password1, password2)
}