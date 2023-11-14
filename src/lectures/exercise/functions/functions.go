//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello", name)
}

func someMessage() string {
	return "This is a super secret message!"
}

func sum3(a int, b int, c int) int {
	return a + b + c
}

func someNumber() int {
	return 235
}

func someNumbers2() [2]int {
	return [2]int{85, 28}
}

func main() {
	hello("Bob")
	fmt.Println(someMessage())

	fmt.Println(sum3(1, 3, 5))
	fmt.Println(someNumber())
	fmt.Println(someNumbers2())
	fmt.Println(sum3(someNumber(), someNumbers2()[0], someNumbers2()[1]))
}
