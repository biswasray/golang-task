package main

import "fmt"

func main() {
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)

	fmt.Println("true && true = ", (true && true))
	fmt.Println("true && false = ", (true && false))
	fmt.Println("false && true = ", (false && true))
	fmt.Println("false && false = ", (false && false))

	fmt.Println("true || true = ", (true || true))
	fmt.Println("true || false = ", (true || false))
	fmt.Println("false || true = ", (false || true))
	fmt.Println("false || false = ", (false || false))

	fmt.Println("(9 == 9) && (9 < 98)", (9 == 9) && (9 < 98))
	fmt.Println("!(4 < 0)", !(4 < 0))
	fmt.Println("!(9 < 98) || (7 >= 98)", !(9 < 98) || (7 >= 98))
	fmt.Println("!((4 > 0) && !(9 > 98))", !((4 > 0) && !(9 > 98)))
	fmt.Println("(4 <= 0) || (9 <= 98)", (4 <= 0) || (9 <= 98))
	fmt.Println("!(9 >= 98) && (4 >= 0) || (9 >= 9)", !(9 >= 98) && (4 >= 0) || (9 >= 9))
}
