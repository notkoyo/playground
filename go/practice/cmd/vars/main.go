package main

import (
	"fmt"
	"unicode/utf8"
)

func foo() string {
	bar := "bar"
	return bar
}

func main() {
	// int8, int16, int32, int64 - specifies maximum bits
	// uint (unsigned int - only stores positive integers)
	var intNum int = 69
	fmt.Println(intNum)

	// float32, float64 - float64 more precise but uses more memory
	var floatNum float64 = 69.69
	fmt.Println(floatNum)

	var myString string = "Hello World"
	fmt.Println(myString)

	// number of bytes, not length of string
	fmt.Println(len("test"))

	// test length of string
	fmt.Println(utf8.RuneCountInString("Hello World"))

	// Must use ' ' instead of " " for runes
	var myRune rune = 'a'
	fmt.Println(myRune)

	// True or false
	var myBool bool = true
	fmt.Println(myBool)

	/*
		Go sets default values for unset variables:
		- For all unsigned ints, ints, floats and runes the default value is 0
		- For strings the default value is ' '
		- For bools the default value is false
	*/

	// You can ommit the type if the value is asigned from the start as it is inferred
	var helloWorld = "Hello World"
	fmt.Println(helloWorld)

	// The shorthand for helloWorld variable
	helloWorld2 := "Hello World!"
	fmt.Println(helloWorld2)

	// You can also initialize multiple variables at once
	num1, num2, num3 := 1, 2, 3
	fmt.Println(num1, num2, num3)

	// Always assign the type when it is not obvious
	var notFunc string = foo()
	fmt.Println(notFunc)

	// Const cant be changed once created and must be initialized
	// const myConst string - invalid syntax
	const myConst string = "const value"
	fmt.Println(myConst)
}
