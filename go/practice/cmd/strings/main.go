package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"H", "e", "l", "l", "o", ", ", "W", "o", "r", "l", "d", "!"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr string = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)
}
