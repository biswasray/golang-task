package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter")
	sc := bufio.NewReader(os.Stdin)
	// data0, _ := sc.ReadString('\n')
	// fmt.Printf("%s %T", data0, data0)
	// data1, _ := sc.ReadByte()
	// fmt.Printf("%c %T", data1, data1)
	data2, size, _ := sc.ReadRune()
	fmt.Printf("%c %T %d", data2, data2, size)
}
