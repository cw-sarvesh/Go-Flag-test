package main

import (
	"fmt"
)

func main() {
	var a int = 20
	var x *int

	x = &a

	fmt.Printf("Address of the a is %x\n", x)
	fmt.Printf("Value of *x is %d\n", *x)
}
