package main

import "fmt"

func main() {
	var x int = 9
	p := &x
	fmt.Println(*p)
}
