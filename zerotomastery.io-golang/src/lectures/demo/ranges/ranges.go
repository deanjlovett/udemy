package main

import "fmt"

func main() {
	slice := []string {"Hello","world","!"}

	// will not work because runes in string 
	// are encoded using utf-8
	//
	// for i := 0; i<10; i++ {
	// 	slice[0][i]
	// }

	for i, element := range slice {
		fmt.Print(i, " ", element, " : ")
		for _, ch := range element {
			fmt.Printf(" %q", ch)
		}
		fmt.Println()
	}
}
