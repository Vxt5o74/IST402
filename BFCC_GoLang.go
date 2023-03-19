package main

import (
	"bufio"
	"fmt"
	"os"
)

// Encrypt a string and return ciphertext
func encrypt(str string, shift int) string {
	var result string

	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			result += string('a' + (letter-'a'+rune(shift))%26)
		} else if letter >= 'A' && letter <= 'Z' {
			result += string('A' + (letter-'A'+rune(shift))%26)
		} else {
			result += string(letter)
		}
	}

	return result
}

// Decrypt a string and return plaintext
func decrypt(str string, shift int) string {
	var result string

	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			result += string('a' + (letter-'a'-rune(shift)+26)%26)
		} else if letter >= 'A' && letter <= 'Z' {
			result += string('A' + (letter-'A'-rune(shift)+26)%26)
		} else {
			result += string(letter)
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter a message: ")
	scanner.Scan()
	message := scanner.Text()

	shift := 3

	ciphertext := encrypt(message, shift)
	plaintext := decrypt(ciphertext, shift)

	fmt.Println("\nCiphertext: ", ciphertext)
	fmt.Println("Shift: ", shift)
	fmt.Println("Plaintext: ", plaintext)

}
