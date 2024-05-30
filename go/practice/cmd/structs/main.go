package main

import (
	"fmt"
)

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
	age  int
	job  string
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println(("You can make it there!"))
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

func main() {
	var me owner = owner{"Kaiden", 23, "Software Developer"}
	var myEngine gasEngine = gasEngine{21, 15, me}
	fmt.Printf("MPG: %v\nGallons: %v\nOwner Info:\n Name: %v \n Age: %v \n Job: %v\n", myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.ownerInfo.age, myEngine.ownerInfo.job)

	// Anonymous Struct
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Printf("\n%v, %v\n", myEngine2.mpg, myEngine2.gallons)

	// Calling methods
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	// Using interfaces
	myElectricEngine := electricEngine{100, 50}
	canMakeIt(myEngine, 50)
	canMakeIt(myElectricEngine, 50)
}
