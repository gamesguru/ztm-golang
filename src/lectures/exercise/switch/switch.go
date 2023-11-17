//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age_check(age int) {
	fmt.Print("Age ", age, " -> ")
	switch {
	case age == 0:
		fmt.Println("newborn")
	case age < 4:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	case age >= 18:
		fmt.Println("adult")
	default:
		fmt.Println("UNKNOWN")
	}
}

func main() {
	for i := 0; i < 40; i++ {
		age_check(i)
	}
}
