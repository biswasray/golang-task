package main

import "fmt"

func main() {
	a, b := 60, 13
	a += b
	fmt.Println("a =", a)
	b -= 10
	fmt.Println("b =", b)
	a *= 2
	fmt.Println("a =", a)
	a /= 10
	fmt.Println("a =", a)
	a %= 5
	fmt.Println("a =", a)
	b <<= 2
	fmt.Println("b =", b)
	b >>= 10
	fmt.Println("b =", b)
	a &= b
	fmt.Println("a =", a)
	b ^= 10
	fmt.Println("b =", b)
	a |= 2
	fmt.Println("a =", a)
}
