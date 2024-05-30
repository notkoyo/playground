package main

import (
	"fmt"
	"time"
)

func main() {
	var intArr [3]int32

	// Print array values
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Print array memory locations (uses &)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Assign values
	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr2[0])
	fmt.Println(intArr2[1:3])
	fmt.Println(intArr2)

	// Slices are arrays with additional functionality
	intSlice := []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // Appends 7 to the end of the array
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// Append multiple values at once
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Use the make function/method instead
	// make(type[], len, cap)
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// Maps are similar to maps in other languages
	// map [keys] values
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 46}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])

	// delete function
	delete(myMap2, "Adam")

	// Map returns <optional> second value of type bool
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("\nPerson exists and is %v years old!", age)
	} else {
		fmt.Printf("Person does not exist!\n")
	}

	// Iterating through maps
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// Go while loop
	x := 0
	for x <= 10 {
		fmt.Println(x)
		x++
	}

	// Go For loop similar to For loop in JS
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	speed()
}

func speed() {
	var n int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without set capacity: %v\n", timeLoop(testSlice1, n))
	fmt.Printf("Total time with set capacity: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
