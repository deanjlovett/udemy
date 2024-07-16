package display

import "fmt"

// visibility of functions in a package 
// is governed by the capitalization of the 
// first letter in the name
//
// UPPERCASE is visibile or public
//
// lowercase is hidden or private 
//
//
// since the "D" in Display is capitalized
// this will be public
// this will be visible outside the package

func Display(msg string) {
	fmt.Println(msg)
}
