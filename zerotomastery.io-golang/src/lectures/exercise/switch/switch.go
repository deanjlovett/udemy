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

func main() {

	var catagory string

	switch age := 14; {
	case age == 0: 
		catagory = "newborn"
	case age <= 3: 
		catagory = "toddler"
	case age <= 12: 
		catagory = "child"
	case age <= 17:
		catagory = "teenager"
	default:
		catagory = "adult"
	}
	fmt.Println("cat-a-gory is", catagory)
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+
}
