package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("max int8 =", math.MaxInt8)
	fmt.Println("min int8 =", math.MinInt8)

	fmt.Println("max int16 =", math.MaxInt16)
	fmt.Println("min int16 =", math.MinInt16)

	fmt.Println("max int32 | rune =", math.MaxInt32)
	fmt.Println("min int32 | rune =", math.MinInt32)

	fmt.Println("max int64 =", math.MaxInt64)
	fmt.Println("max int64 =", math.MinInt64)

	fmt.Println("max uint8 | byte =", math.MaxUint8)
	fmt.Println("max uint16 =", math.MaxUint16)
	fmt.Println("max uint32 =", math.MaxUint32)

	var UInt64 uint64 = math.MaxUint64
	fmt.Println("max uint64 =", UInt64)

	var UInt uint = 0
	fmt.Println("max uint =", UInt-1)
	var Int int = 0
	fmt.Println("max int =", int((UInt-1)/2))
	fmt.Println("min int =", Int-int((UInt-1)/2+1))

	// uintptr
	// NaN
	// float32
	// float64
	// complex64
	// complex128

	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Printf("1 + 1 = %.1f \n", 1.0+1.0)

	for i := .0; ; i += .2 {
		if i == 10. {
			fmt.Printf("i == 10 OK!, i - %.20f\n", i)
			break
		} else if i > 10. {
			fmt.Printf("i > 10 ERROR, i = %.20f\n", i) // 10.19999999999999573674
			break
		}
	}

	for i := .0; ; i += .2 {
		if math.Abs(10-i) < .2/2 {
			fmt.Printf("i == 10 OK!, i - %.20f\n", i) // 9.99999999999999644729
			break
		} else if i > 10. {
			fmt.Printf("i > 10 ERROR, i = %.20f\n", i)
			break
		}
	}

	for i := .0; ; i += .2 {
		if i <= 10.+.2/2 && i >= 10.-.2/2 {
			fmt.Printf("i == 10 OK!, i - %.20f\n", i) // 9.99999999999999644729
			break
		} else if i > 10. {
			fmt.Printf("i > 10 ERROR, i = %.20f\n", i)
			break
		}
	}

	var text string = "Hello \t World!"
	fmt.Println("text = \"", text, "\"")

	text = `Hello \t World!`
	fmt.Println("text = `" + text + "`")

	fmt.Println("text[3] =", string(text[3]))
	fmt.Println("len =", len(text))

	// true
	// false

}
