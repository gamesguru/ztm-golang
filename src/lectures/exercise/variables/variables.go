//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
    fmt.Println("Hello!")

    var favColor = "Blue";

    var birthYear, age = 1992, 30
    var (
        initialFirst string = "S"
        initialLast string = "J"
    )

    var ageDays int
    ageDays = 365 * age

    // Print info
    // fmt.Sprintf("%s. %s. (%s, age: %s, days old: %s, fav color: %s)", initialFirst, initialLast, age, birthYear, ageDays, favColor)

    fmt.Println(initialFirst)
    fmt.Println(initialLast)

    fmt.Println(favColor)
    fmt.Println(birthYear)
    fmt.Println(age)
    fmt.Println(ageDays)

    fmt.Println("Done!")
}
