package main

import "fmt"

func main() {

	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello!"
	fmt.Println(y)

	fmt.Println(x == y)

	var b = "Hello boys"
	fmt.Println(b)

	c := 12
	fmt.Println(c)

	// Область видимости

	// const
	const ConstA = 100
	// const ConstB := 200 ! ERROR
	const ConstC int8 = 20

	var (
		A = 5
		B = 10
		C = 15
	)

	fmt.Println(A, B, C)

	const (
		q = 88
	)

	// const (
	// 	 q = 5 ! ERROR
	// )

	fmt.Println(q)

	fmt.Print("Enter a number: ")
	var input float64 = 100.123
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println("Input Error")
	}
	fmt.Println(input)

}
