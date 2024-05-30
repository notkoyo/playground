package main

import "fmt"

func main() {
	// Assign new memory address to pointer
	var p *int32 = new(int32)
	var i int32 = 10

	p = &i
	*p = 1

	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)
}
