package main

import "fmt"

func main() {

	// printing

	// fmt.Println("Hello World!")
	// fmt.Println("Go" + "lang")
	// fmt.Println("1 + 4 = ", 1 + 4)
	// fmt.Println("false", true && false)
	// fmt.Println("true", true || false)
	// fmt.Println(!true)

	// variable declaration + shorthand declaration and initialization

	// var num1 = 1
	// fmt.Println("1", num1)

	// var num2, num3 int = 2, 3
	// fmt.Println("2 3 === ", num2, num3)

	// var zero int
	// fmt.Println("0", zero)

	// var false bool
	// fmt.Println("false?", false)

	// var emptyString string
	// fmt.Println("0", len(emptyString))

	// lastVar := "last variable"

	// fmt.Println(lastVar)

	// lastVar = "new value to prove a point lol"

	// fmt.Println(lastVar)

	// for loops

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}
}
