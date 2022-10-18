package main

import "fmt"

func main() {
	a, b := 36, 4

	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b
	fmt.Printf("sum=%d\nsub=%d\nmul=%d\ndiv=%d\nmod=%d", c, d, e, f, g)

	a++
	fmt.Println("Increament of 'a' =", a)
	b--
	fmt.Println("Increament of 'b' =", b)

}
