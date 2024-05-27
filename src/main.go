package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	// var array = [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	// fmt.Println(sumOfEvenArray(array))

	/*

		1. **Create a Map and Add Key-Value Pairs:**
		   Initialize an empty map called `ages` where the keys are names (strings) and the values are ages (integers). Add a few key-value pairs to the map and print the map.

		2. **Retrieve Values from a Map:**
		   Use the `ages` map from the previous problem. Retrieve and print the age of a specific person (key) from the map.

		3. **Check if a Key Exists:**
		   Use the `ages` map again. Check if a specific name exists as a key in the map. Print "Exists" if the name is found and "Doesn't Exist" otherwise.

		4. **Remove a Key-Value Pair:**
		   Use the `ages` map. Remove a specific key-value pair from the map and print the updated map.

		5. **Count the Number of Key-Value Pairs:**
		   Use the `ages` map. Count and print the number of key-value pairs in the map using the `len` function.

		6. **Clear a Map:**
		   Use the `ages` map. Clear all key-value pairs from the map and print the map to confirm it's empty.

		7. **Declare and Initialize a Map in One Line:**
		   Declare and initialize a new map called `grades` in a single line where the keys are subjects (strings) and the values are grades (floats). Print the map.

		8. **Check Map Equality:**
		   Create two maps `m1` and `m2` with the same key-value pairs. Use the `maps.Equal` function to check if `m1` is equal to `m2` and print the result.

	*/

	// // 1

	// ages := make(map[string]int)

	// ages["Dylan"] = 25
	// ages["Chris"] = 27
	// ages["James"] = 28

	// fmt.Println("ages map", ages)

	// // 2

	// fmt.Println("Dylan is", ages["Dylan"], "years old.")
	// fmt.Println("Should print 0 since key doesn't exist", ages["Irais"])

	// // 3

	// _, prs := ages["Dylan"]

	// if prs {
	//   fmt.Println("Exists")
	// } else {
	//   fmt.Println("Doesn't Exist")
	// }

	// //4
	// delete(ages, "Chris")
	// fmt.Println("deleted", ages)

	// //5
	// fmt.Println("2?", len(ages))

	// //6
	// clear(ages)

	// fmt.Println("should be empty", ages)

	// // 7
	// grades := map[string]float64{"Comp Sci": 97.5, "Psychology": 98.5, "Entomology": 96.5}
	// students := map[string]int{"Chris": 12, "Dylan": 11}

	// fmt.Println("grades: ", grades)
	// fmt.Println("students", students)

	// letterFrequency("Hello")
	// fmt.Println("GRADES", makeGradesMap())
	// cart := map[string]float64{"Apples": 0.57, "Olipop": 1.50, "Bananas": 1.00, "Kefir": 3.47}

	// addToShoppingCart(cart)
	// removeCartItem(cart, "Apples")
	// fmt.Println("Your grand total is: ", calculateCartTotal(cart))
	// fmt.Print("You're updated cart: ", cart)
	// fmt.Println("Should return 3", rangePractice())
	// fmt.Println(moreRangePractice())
	// fmt.Println(stringRangePractice("West Hollywood Library"))
	// fmt.Println(multipleReturnsFunc("West", "Hollywood"))
	// fmt.Println(myVariadicFunc(9, 9, 9))
	// closures

	// result := intSeq()

	// fmt.Println(result())
	// fmt.Println(result())
	// fmt.Println(result())
	// fmt.Println("Should be 4", result())
	// fmt.Println("should be 4", lengthOfLastWord("   fly me   to   the moon  "))
	// fmt.Println("should be 5", lengthOfLastWord("hello world"))
	// fmt.Println(isPalindrome(121))
	// fmt.Println(isPalindrome(123))
	// arr := []int{1, 3, 4, 5, 7}
	// fmt.Println("should be 4", binarySearchAlgo(arr, 7))
	// fmt.Println("should be false", containsDuplicate(arr))
	// fmt.Println("should be false", isAnagram("Meltt", "Melt"))
	// fmt.Println("should be true", isAnagram("Dylan Gardner", "rendraG nalyD"))
	// myStr := "Hello"
	// fmt.Println(iterateMyStr(myStr))

	// value := 8
	// fmt.Println("init value", value)
	// regularFunction(value)
	// fmt.Println("value after call to regularFunc should be 8", value)
	// pointerFuncExample(&value)
	// fmt.Println("should be 10", value)
	// fmt.Println("should be memory address", &value)

	// digits := []int{1, 2, 3, 4}
	// fmt.Print("should be [1, 2, 3, 5]", plusOne(digits))
	// potentialPairs := []string{"cd", "ac", "dc", "ca", "zz", "lh"}
	// fmt.Println("should be 2", maximumNumberOfStringPairs(potentialPairs))

	// names := []string{"Mary", "John", "Emma"}
	//  heights := []int{180, 165, 170}
	// fmt.Println(sortPeople(names, heights))
	// fmt.Println(createRectangle(12.0, 12.0))

	makeBook("The How of Happiness", "Sonja Lyubomirsky", 384, 2010)
	makeBook("Educated", "Tara Something", 400, 2013)
	makeBook("Homebody", "Rupi Kaur", 188, 2020)
}

/*
	  Title (string)
      Author (string)
      Pages (int)
      PublishedYear (int)

*/

/*
  STRUCTS

    Create a struct named Rectangle with fields length and width, both of type float64. Define two methods for the Rectangle struct:

	area() method that calculates and returns the area of the rectangle.
	perimeter() method that calculates and returns the perimeter of the rectangle.
	Then, write a main function to create a Rectangle instance, initialize its fields, and call both the area() and perimeter() methods to display the results.
*/

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

func (r *rectangle) perimeter() float64 {
	return 2*r.width + 2*r.length
}

func createRectangle(length float64, width float64) rectangle {
	newRectangle := rectangle{length: length, width: width}
	newRectangle.length = newRectangle.area()
	newRectangle.width = newRectangle.perimeter()
	return newRectangle

}

/*

Struct Definition
First, define a struct to represent a book. Each book should have the following fields:

Title (string)
Author (string)
Pages (int)
PublishedYear (int)

Functions to Implement
UpdateBook: Updates the details of a book.
DeleteBook: Deletes a book from the collection.

*/

type book struct {
	title         string
	author        string
	pages         int
	publishedYear int
}

var collection = []book{}

// constructor function 
func makeBook(title string, author string, pages int, publishedYear int) book {
	newBook := book{title: title, author: author, pages: pages, publishedYear: publishedYear}
	collection = newBook.addBook(collection)
	newBook.listBook(collection)
	book := newBook.findBookByTitle(collection, "Homebody")
	fmt.Println(book)


	return newBook
}

// AddBook: Adds a new book to the collection.
func (b book) addBook(collection []book) []book {
	return append(collection, b)
}
// ListBooks: Lists all books in the collection.
func (b book) listBook (collection []book) []book {
  return collection
}

// FindBookByTitle: Finds a book by its title.
func (b book) findBookByTitle(collection []book, title string) book {
  var response book
  for i:= 0; i < len(collection); i++ {
    if collection[i].title == title {
	  response = collection[i]
	  break
	} 
  }
  return response
}

// **ALGO 1**

/*
  Write a program that takes an array of integers as input and calculates the sum of all the even numbers in the array.
*/

// func sumOfEvenArray(numsArray [14]int) int {
// 	result := 0
// 	for i := 0; i < len(numsArray); i++ {
// 		if numsArray[i]%2 == 0 {
// 			result += numsArray[i]
// 		}
// 	}
// 	return result
// }

// **ALGO 2**

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

// **ALGO 3**

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

// Algo 4
/*
Counting Letters:
  Write a program that takes a string as input and counts the frequency of
  each letter in the string. Store the letter frequencies in a map and print
  the map.
*/

/*
  make a map using make method

  use a loop to go thru string

  at each iteration check if key exists uf it does
  update key frequency aka add 1 (use prs)

  if not make it an entity add the letter as key and value as 1

  print the map
*/

// func letterFrequency(str string) {
// 	hashMap := make(map[string]int)

// 	for i := 0; i < len(str); i++ {
// 		currentChar := string(str[i])
// 		_, prs := hashMap[currentChar]
// 		if prs {
// 			hashMap[currentChar] += 1
// 		} else {
// 			hashMap[currentChar] = 1
// 		}
// 	}
// 	fmt.Println("Populated map: ", hashMap)
// }

// ALGO 5
/*
Student Grades:

Create a program that stores the grades of students in different
subjects. Use a map where the keys are student names (strings)
and the values are maps of subject grades (strings) to numerical
scores (floats). Add grades for at least three students in
different subjects and print the overall grade for each student.

*/

// func makeGradesMap() map[string]map[string]float64 {
// 	grades := map[string]map[string]float64{
// 		"Dylan": {"Psychology": 98.5, "Art": 100.0},
// 		"Rambo": {"Music": 100.0, "Psychology": 98.5},
// 		"Chris": {"Music": 98.5, "Art": 100.0},
// 	}
// 	return grades
// }

// ALGO 6
/*
4. **Shopping Cart:**
   Create a program that simulates a shopping cart. Use a map to
   store items (strings) as keys and their prices (floats) as
   values. Allow the user to add items to the cart, remove items,
   and calculate the total price of all items in the cart.
*/

/*
  make a map using make since i dont know what i'll be adding yet

  make the items be the keys -will be strings anf the prices be
  values  - will be float64s

  use bufio to create a reader and osStdin to read terminal input
    write logic to add that to new item to map
	return remaining items

  make another func to recieve a item to remove and remove it then
  return remaining items

  make another function to look through values in map and cal
  total of all items
  return total


*/

func addToShoppingCart(cart map[string]float64) map[string]float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter item name and price omit the $ pls: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	floatVal, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
	} else {
		cart[parts[0]] = floatVal
	}

	return cart
}

func removeCartItem(cart map[string]float64, itemToRemove string) map[string]float64 {
	delete(cart, itemToRemove)

	return cart
}

func truncateToTwoDecimalPlaces(num float64) float64 {
	return math.Trunc(num*100) / 100
}

func calculateCartTotal(cart map[string]float64) interface{} {

	var total float64 = 0.00

	if len(cart) > 0 {
		for _, value := range cart {
			total += value
		}
		return truncateToTwoDecimalPlaces(total)
	} else {
		return "Shopping cart is empty."
	}

}

// var array = [5]int{1, 2, 3, 4, 5}

// func rangePractice() int {
//   var answer int
//   for _, elem := range array {
// 	if elem == 3 {
// 		answer = elem
// 	}
//   }
//  return answer
// }

// make a func that iterates over an map and using range and push all the keys into and array and all the values return both arrays

func moreRangePractice() ([]string, []int) {

	namesAndAgesMap := map[string]int{"Yuri": 8, "Bingo": 1, "Lassie": 1, "Cokita": 3}

	// initializing slices
	keyArray := make([]string, 0, 10)
	valueArray := make([]int, 0, 10)

	for key, value := range namesAndAgesMap {
		keyArray = append(keyArray, key)
		valueArray = append(valueArray, value)
	}

	return keyArray, valueArray

}

// make a func that iterates over a string using range and create a map that stores index of letters as key and letter as value of that key. Return the map

func stringRangePractice(stringToIterateOver string) map[int]string {
	stringRecord := make(map[int]string)

	for i, elem := range stringToIterateOver {
		if string(elem) != " " {
			stringRecord[i] = string(elem)

		}
	}
	return stringRecord
}

// make a func that returns multiple strings

func multipleReturnsFunc(string1 string, string2 string) (string, string) {
	return string1, string2
}

// make a variadic function
func myVariadicFunc(nums ...int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// time and space: o(n)
func lengthOfLastWord(s string) int {
	arr := strings.Fields(s)
	lastElemPosition := len(arr) - 1
	result := len(arr[lastElemPosition])

	return result
}

/*
  APPROACH

  make a array remove the white spaces
  return the last elem
*/

func isPalindrome(x int) bool {
	numStr := strconv.Itoa(x)
	arr := []string{}

	for i := len(numStr) - 1; i >= 0; i-- {
		arr = append(arr, string(numStr[i]))
	}

	arrStr := strings.Join(arr, "")

	return numStr == arrStr
}

/*
	APPROACH
	remember that x is an int and mine will be strings
	turn into string
	iterate through string start at the end append strings to an array
	compare the number turned into a string and array turned into string
	if they are true return true else false
*/

func binarySearchAlgo(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		midIdx := (left + right) / 2
		if nums[midIdx] == target {
			return midIdx
		} else if target < nums[midIdx] {
			right = midIdx - 1
		} else {
			left = midIdx + 1
		}

	}

	return -1
}

// o(n) space and time
func containsDuplicate(nums []int) bool {
	frequencyMap := make(map[int]int)

	for _, value := range nums {

		if _, key := frequencyMap[value]; key {
			return true
		} else if !key {
			frequencyMap[value] = 1
		}
	}
	return false
}

/*

	make a empty map (map will hold frequency of ints)
	make a loop to go through elements and create map at the same time
	if we need to increment frequency for any at anytime we will return true

	if we make it out of loop then we didnt find any dups so return false

*/

// func isAnagram(s string, t string) bool {
// 	sMap := make(map[string]int)

// 	if len(s) != len(t) {
// 		return false
// 	}

// 	for i := 0; i < len(s); i++ {
// 		curChar := string(s[i])
// 		if _, key := sMap[curChar]; key {
// 			sMap[curChar] += 1
// 		} else if !key {
// 			sMap[curChar] = 1
// 		}
// 	}

// 	for i := 0; i < len(t); i++ {
// 		curChar := string(t[i])
// 		if _, key := sMap[curChar]; key {
// 			sMap[curChar] -= 1
// 		}
// 	}

// 	var flag bool

// 	for _, value := range sMap {
// 		if value == 0 {
// 			flag = true
// 		} else if value > 0 {
// 			return false
// 		}
// 	}

// 	return flag
// }

// func iterateMyStr(str string) bool {

//   // iterates through bytes
//   for i := 0; i < len(str); i++ {
//     fmt.Println("byte", str[i])
//   }

//   // iterates through runes
//   for _, value := range str {
// 	fmt.Println("rune", value)
//   }
//   return true
// }

// pointers in GO
func regularFunction(val int) {
	val = 0
}

func pointerFuncExample(val *int) {
	*val += 2
}

func plusOne(digits []int) []int {
	result := []int{1}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}

	result = append(result, digits...)

	return result

}

/*
 make a loop to start at the end of digits
 when we come across a 9
 turn it into a zero

 else if the number is not nine
 then increment by one and return digits


 outisde return a slice with one in the beginning for the instance where we have only one 9
*/

func maximumNumberOfStringPairs(words []string) int {
	pairsMap := make(map[string]string)
	totalPairs := 0

	for i := 0; i < len(words); i++ {
		_, key := pairsMap[reverseCurrentWord(string(words[i]))]
		if key {
			totalPairs += 1
		} else {
			pairsMap[string(words[i])] = string(words[i])
		}
	}

	return totalPairs
}

func reverseCurrentWord(word string) string {
	result := []string{}

	for i := len(word) - 1; i >= 0; i-- {
		result = append(result, string(word[i]))
	}

	return strings.Join(result, "")

}

func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] < heights[j] {
				tempForHeights := heights[i]
				tempForNames := names[i]
				heights[i] = heights[j]
				heights[j] = tempForHeights
				names[i] = names[j]
				names[j] = tempForNames
			}
		}
	}
	return names
}
