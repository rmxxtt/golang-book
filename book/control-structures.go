package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for i = 1; i <= 10; i++ {
		fmt.Println(i)
	}

	if i%2 == 0 {
		// even
	} else {
		// odd
	}

	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Четное")
		} else {
			fmt.Println(i, "Нечетный")
		}
	}

	name := "John"
	switch name {
	case "Ellis":
		fmt.Println("Hello", name)
	case "John":
		fmt.Println("Hello", name)
		//break Не надо!)
	case "Alister":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Who are you?")
	}

	// Задача 1
	// Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).

	for i = 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(" ", i)
		}
	}

	fmt.Println()
	// Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно
	// вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём,
	// так и пяти — «FizzBuzz».

	for i = 1; i <= 100; i++ {
		switch { // выражение не требуется!
		case i%3 == 0 && i%5 == 0:
			fmt.Print(" Fizz-Buzz")
		case i%3 == 0:
			fmt.Print(" Fizz")
		case i%5 == 0:
			fmt.Print(" Buzz")
		default:
			fmt.Print(" ", i)
		}
	}

}
