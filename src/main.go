package main

import "fmt"

func main() {

	// printing

	fmt.Println("Hello World!")
	fmt.Println("Go" + "lang")
	fmt.Println("1 + 4 = ", 1 + 4)
	fmt.Println("false", true && false)
	fmt.Println("true", true || false)
	fmt.Println(!true)

	// variable declaration + shorthand declaration and initialization

	var num1 = 1
	fmt.Println("1", num1)

	var num2, num3 int = 2, 3
	fmt.Println("2 3 === ", num2, num3)

	var zero int
	fmt.Println("0", zero)

	var false bool
	fmt.Println("false?", false)

	var emptyString string
	fmt.Println("0", len(emptyString))

	lastVar := "last variable"

	fmt.Println(lastVar)

	lastVar = "new value to prove a point lol"

	fmt.Println(lastVar)

	// for loops

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	// if else chains
	if num := 9; num < 0 {
		fmt.Println(num, "number is negative")
	} else if num < 10 {
		fmt.Println(num, "number has 1 digit")
	} else {
		fmt.Println(num, "number has multiple digits")
	}

	// all three lessons in one
	const two = 2.0

	for i := 0; i <= 5; i++ {
		if float64(i)/two == 2.0 {
			fmt.Println("should be 4", i)
		}
	}

  // switches
	const day = 4

	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	default:
		fmt.Println("Today is Sunday")
	}

  // arrays
  var array [6]int

  fmt.Println("initial array", array)

  count:= 1

  for i:= 0; i < len(array); i++ {
    array[i] = count
    count += 1
  }

  fmt.Println("after assigning elems", array)

  array2:= [3]int{1, 2, 3}

  fmt.Println("array2", array2)
    
}
