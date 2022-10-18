package main

import "fmt"

func main() {
	fmt.Println("9 == 9", 9 == 9)
	fmt.Println("4 == '4'", 4 == '4')

	fmt.Println("4 != 0", 4 != 0)
	fmt.Println("9 != 9", 9 != 9)

	fmt.Println("4 < 0", 4 < 0)
	fmt.Println("9 < 98", 9 < 98)

	fmt.Println("4 > 0", 4 > 0)
	fmt.Println("9 > 98", 9 > 98)

	fmt.Println("4 <= 0", 4 <= 0)
	fmt.Println("9 <= 9", 9 <= 9)
	fmt.Println("9 <= 98", 9 <= 98)

	fmt.Println("4 >= 0", 4 >= 0)
	fmt.Println("9 >= 9", 9 >= 9)
	fmt.Println("9 >= 98", 9 >= 98)

}
