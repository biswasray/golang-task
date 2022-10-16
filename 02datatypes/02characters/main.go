package main

import "fmt"

func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'
	/*
	 * In Go, there is no char data type. It uses byte and rune to represent character values.
	 * The byte data type represents ASCII characters while the rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format.
	 */

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}
