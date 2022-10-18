package main

import "fmt"

func main() {
	a, b := 60, 13
	fmt.Println("a&b", a&b)
	fmt.Println("a|b", a|b)
	fmt.Println("a^b", a^b)
	fmt.Println("a>>2", a>>2)
	fmt.Println("a<<2", a<<2)

	fmt.Println("^a", ^a)
}
