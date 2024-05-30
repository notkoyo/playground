package main

import "fmt"

// Generics can be used with structs to enable multiple types

type gasEngine struct {
	mpg     float32
	gallons float32
}

type electricEngine struct {
	mpkwh float32
	kwh   float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var intSlice = []int{1, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 1}
	fmt.Println(sumSlice[float32](float32Slice))

	var float64Slice = []float64{0, 0}
	fmt.Println(sumSlice[float64](float64Slice))

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			mpg:     40,
			gallons: 12.4,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			mpkwh: 100,
			kwh:   1000,
		},
	}

	fmt.Println(gasCar, electricCar)

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
