package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {

	// printing

	// 	fmt.Println("Hello World!")
	// 	fmt.Println("Go" + "lang")
	// 	fmt.Println("1 + 4 = ", 1 + 4)
	// 	fmt.Println("false", true && false)
	// 	fmt.Println("true", true || false)
	// 	fmt.Println(!true)

	// 	// variable declaration + shorthand declaration and initialization

	// 	var num1 = 1
	// 	fmt.Println("1", num1)

	// 	var num2, num3 int = 2, 3
	// 	fmt.Println("2 3 === ", num2, num3)

	// 	var zero int
	// 	fmt.Println("0", zero)

	// 	var false bool
	// 	fmt.Println("false?", false)
	// 	fmt.Printf("This variable is a %T\n ", false)

	// 	var emptyString string = "This is go!"
	// 	fmt.Println("0", len(emptyString))
	// 	fmt.Printf("This variable is a %T\n" , emptyString)

	// 	lastVar := "last variable"

	// 	fmt.Println(lastVar)

	// 	lastVar = "new value to prove a point lol"

	// 	fmt.Println(lastVar)

	// 	// for loops

	// 	for i := 0; i <= 5; i++ {
	// 		fmt.Println(i)
	// 	}

	// 	for i := range 3 {
	// 		fmt.Println("range", i)
	// 	}

	// 	// if else chains
	// 	if num := 9; num < 0 {
	// 		fmt.Println(num, "number is negative")
	// 	} else if num < 10 {
	// 		fmt.Println(num, "number has 1 digit")
	// 	} else {
	// 		fmt.Println(num, "number has multiple digits")
	// 	}

	// 	// all three lessons in one
	// 	const two = 2.0

	// 	for i := 0; i <= 5; i++ {
	// 		if float64(i)/two == 2.0 {
	// 			fmt.Println("should be 4", i)
	// 		}
	// 	}

	//   // switches
	// 	const day = 4

	// 	switch day {
	// 	case 1:
	// 		fmt.Println("Today is Monday")
	// 	case 2:
	// 		fmt.Println("Today is Tuesday")
	// 	case 3:
	// 		fmt.Println("Today is Wednesday")
	// 	case 4:
	// 		fmt.Println("Today is Thursday")
	// 	case 5:
	// 		fmt.Println("Today is Friday")
	// 	case 6:
	// 		fmt.Println("Today is Saturday")
	// 	default:
	// 		fmt.Println("Today is Sunday")
	// 	}

	//   // arrays
	//   var array [6]int

	//   fmt.Println("initial array", array)

	//   count:= 1

	//   for i:= 0; i < len(array); i++ {
	//     array[i] = count
	//     count += 1
	//   }

	//   fmt.Println("after assigning elems", array)

	//   array2:= [3]int{1, 2, 3}

	//   fmt.Println("array2", array2)

	// printing input from command line

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text ")

	// input, _ := reader.ReadString('\n')

	// fmt.Println("You entered:", input)

	// fmt.Print("Enter a number: ")
	// numInput, _ := reader.ReadString('\n')

	// // converting str to int
	// aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Value of number: ", aFloat)
	// }

	// math.round rounds to nearest even integer, the output will be the number that is exactly halfway between
	// the two integers

	// math pkg
	// fmt.Println("4?", math.Round(4.5))
	// fmt.Println("8?", math.Round(8.5))
	// fmt.Println("10?", math.Round(9.5))

	// i1, i2, i3 := 12, 45, 68

	// intSum := i1 + i2 + i3

	// fmt.Println("intSum", intSum)

	// f1, f2, f3 := 23.5, 65.1, 76.3

	// floatSum := f1 + f2 + f3

	// fmt.Println("floatSum", floatSum)
	// // safest way in go of rounding a fractional number and holding on to digits after the decimal point
	// floatSum = math.Round(floatSum*100) / 100

	// fmt.Println("This sum is now: ", floatSum)

	// var arg1 string = "45.1  "
	// var arg2 string = "  2.00  "

	// fmt.Println("Result should be 47.1", calculate(arg1, arg2))
	// factorial()

	// in brackets put the length of array, then the types of values, then finally the elements
	var array = [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	fmt.Println(sumOfEvenArray(array))

}

/*
  Write a program that takes an array of integers as input and calculates the sum of all the even numbers in the array.
*/

func sumOfEvenArray(numsArray [14]int) int {
	result := 0
	for i := 0; i < len(numsArray); i++ {
		if numsArray[i]%2 == 0 {
			result += numsArray[i]
		}
	}
	return result
}

// calculate() returns the sum of the two parameters
// func calculate(value1 string, value2 string) float64 {
// 	// Your code goes here.

// 	// Convert the first string to a float64
// 	newVal1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Convert the second string to a float64
// 	newVal2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Calculate and return the result
// 	return newVal1 + newVal2

// }

/*

Write a program in Go that calculates the factorial of a non-negative integer entered by the user. The factorial of a number is the
product of all positive integers less than or equal to that number.

For example, the factorial of 5 (denoted as 5!) is calculated as 5 x 4 x 3 x 2 x 1, which equals 120.
*/

// func factorial() int {
// 	var result int = 1
// 	// assign the reader object to reader variable
// 	// reading data from input source like the terminal or command line
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter a number: ")
// 	numInput, _ := reader.ReadString('\n')

// 	parsedInt, err := strconv.Atoi(strings.TrimSpace(numInput))

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		for i := parsedInt; i > 0; i-- {
// 			result *= i
// 		}
// 	}

// 	fmt.Println("Result: ", result)

// 	return result
// }
