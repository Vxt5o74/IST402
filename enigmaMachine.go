package main

import (
	"fmt"
	"strings"
)

// Setting up rotor structure
type rotor struct {
	wiring string
	offset int
}

// Function for rotors to move forwards in relation to the set mapping/wiring,
// allowing plaintext to be encrypted to ciphertext
func (r *rotor) forward(c byte) byte {
	i := int(c-'A'+byte(r.offset)) % 26
	return r.wiring[i]
}

// Function for rotors to move backwards in relation to the set mapping/wiring,
// allowing ciphertext to then be decrypted back to plaintext
func (r *rotor) backward(c byte) byte {
	i := strings.IndexByte(r.wiring, c)
	i = (i - r.offset + 26) % 26
	return byte(i + 'A')
}

// Setting up reflector structure
type reflector struct {
	wiring string
}

// Function for reflector to map the inputted byte (c) to the letter/value
// that's 13 positions ahead of the given byte (wrapping around after letter/index 26 to avoid unwanted range/value(s))
func (r *reflector) reflect(c byte) byte {
	i := strings.IndexByte(r.wiring, c)
	return r.wiring[(i+13)%26]
}

// Setting up plugboard structure
type plugboard struct {
	pairs map[byte]byte
}

// Function for plugboard to take individual bytes from plaintext
// and checks for a paired value (letter) for the given byte
// and returns that associated value if there is correspondence
func (p *plugboard) plug(c byte) byte {
	if v, ok := p.pairs[c]; ok {
		return v
	}
	return c
}

func (p *plugboard) addConnection(a, b byte) {
	p.pairs[a] = b
	p.pairs[b] = a
}

// Setting up enigma machine structure
type enigmamachine struct {
	rotors    []*rotor
	reflector *reflector
	plugboard *plugboard
}

// Iterates through each element of the plaintext as individual bytes, moves rotor forward, applies
// reflector and plugboard, and then converts "encrypted" byte(s) and returns em as output
func (e *enigmamachine) forward(c byte) byte {
	for i := len(e.rotors) - 1; i >= 0; i-- {
		c = e.rotors[i].forward(c)
	}
	c = e.reflector.reflect(c)
	for i := 0; i < len(e.rotors); i++ {
		c = e.rotors[i].backward(c)
	}
	c = e.plugboard.plug(c)
	return c
}

// Function for moving rotors backwards similar to the function above for moving em forward
func (e *enigmamachine) backward(c byte) byte {
	c = e.plugboard.plug(c)
	for i := len(e.rotors) - 1; i >= 0; i-- {
		c = e.rotors[i].forward(c)
	}
	c = e.reflector.reflect(c)
	for i := 0; i < len(e.rotors); i++ {
		c = e.rotors[i].backward(c)
	}
	return c
}

// Function to encrypt plaintext
func (e *enigmamachine) encrypt(s string) string {
	output := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		output[i] = e.forward(s[i])
	}
	// Performs the forward function above this function, takes the returned bytes
	// and outputs those bytes converted into a string
	return string(output)
}

// Function to decrypt ciphertext
func (e *enigmamachine) decrypt(s string) string {
	output := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		output[i] = e.backward(s[i])
	}
	return string(output)
}

// Function to reset rotor positions to initial values
func (e *enigmamachine) reset() {
	for _, rotor := range e.rotors {
		rotor.offset = 0
	}
}

// Setting up Enigma Machine with predefined rotors, reflector, and plugboard
func main() {
	// Request reflector configuration
	fmt.Println("Please provide a reflector configuration you'd like to use: (a string of every letter in the alphabet used once)")
	var reflector_config string
	fmt.Scanln(&reflector_config)
	
	// Setting up rotor configs
	r1 := &rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", offset: 0}
	r2 := &rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE", offset: 0}
	r3 := &rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO", offset: 0}
	
	// Enters reflector config that the user inputted earlier in after capitalizing entire string
	// (If it wasn't entirely uppercase already)
	ref := &reflector{wiring: strings.ToUpper(reflector_config)}
	pb := &plugboard{pairs: map[byte]byte{}}
	
	// Adding plugboard connections
	pb.addConnection('E', 'Z')
	pb.addConnection('S', 'H')
	pb.addConnection('A', 'F')
	pb.addConnection('G', 'M')
	e = &enigmamachine{rotors: []*rotor{r1, r2, r3}, reflector: ref, plugboard: pb}
	
	// Reset rotor position before encrypting
	e.reset()

	// Establish message to be encrypted and well, encrypt it
	message := "CHEESEBURGER"
	ciphertext := e.encrypt(message)

	// Prints out the original msg for reference and the encrypted ciphertext
	fmt.Println("Message:", message)
	fmt.Println("Ciphertext:", ciphertext)

	// Reset rotor positions after encrypting the message
	e.reset()

	// Decrypts the ciphertext and prints the decrypted plaintext
	plaintext := e.decrypt(ciphertext)
	fmt.Println("Plaintext:", plaintext)
}
