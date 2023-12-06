//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"strings"
	"unicode"
)
	

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	// generic function that takes a closure
	ng := func(str string, fn func(r rune) bool) int {
		count := 0
		for _, char := range str{
			if fn(char) {
				count++
			}
		}
		return count
	}
	//  - Number of letters	
	nl := func(str string) int {
		return ng(str,unicode.IsLetter)
	}


	//  - Number of digits
	nd := func(str string) int {
		return ng(str,unicode.IsDigit)
	}
	// nd := func(str string) int {
	// 	count := 0
	// 	for _, char := range str{
	// 		if unicode.IsDigit(char) {
	// 			count++
	// 		}
	// 	}
	// 	return count
	// }

	//  - Number of spaces
	ns := func(str string) int {
		return ng(str,unicode.IsSpace)
	}
	// ns := func(str string) int {
	// 	count := 0
	// 	for _, char := range str{
	// 		if unicode.IsSpace(char) .IsDigit(char) {
	// 			count++
	// 		}
	// 	}
	// 	return count
	// }

	//  - Number of punctuation marks
	np := func(str string) int {
		return ng(str,unicode.IsPunct)
	}


	fn := func(lns []string) {
		catstr := strings.Join(lns,"")
		//  - Number of letters
		//  - Number of digits
		//  - Number of spaces
		//  - Number of punctuation marks
		fmt.Println()
		fmt.Println("          letters:", nl(catstr))
		fmt.Println("           digits:", nd(catstr))
		fmt.Println("           spaces:", ns(catstr))
		fmt.Println("punctuation marks:", np(catstr))

		fmt.Println()
		fmt.Println("          letters:", ng(catstr,unicode.IsLetter))
		fmt.Println("           digits:", ng(catstr,unicode.IsDigit))
		fmt.Println("           spaces:", ng(catstr,unicode.IsSpace))
		fmt.Println("punctuation marks:", ng(catstr,unicode.IsPunct))

	}
	fn(lines)

}
