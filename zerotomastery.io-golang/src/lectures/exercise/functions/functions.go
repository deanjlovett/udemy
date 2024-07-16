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

//--Requirements:
//* Write a function that accepts a person's name as a 
//  function parameter and displays a greeting to that person.

func greetPerson(name string) {
	fmt.Println("Greetings ",name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func aMessage() string {
	return " >>> any message <<< "
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers(x,y,z int) int {
	return x+y+z
}
//* Write a function that returns any number
func anyNumber() int {
	return 14
}

//* Write a function that returns any two numbers
func anyTwoNumbers(a,b int) (int,int) {
	return a,b
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result

//* Call every function at least once

func main() {

	//--Requirements:
	//* Write a function that accepts a person's name as a 
	//  function parameter and displays a greeting to that person.

	fmt.Println()
	fmt.Println("//--Requirements:")
	fmt.Println("//* Write a function that accepts a person's name as a function")
	fmt.Println("//  parameter and displays a greeting to that person.")
	fmt.Println()
	fmt.Println("Calling greetPerson(\"Fred\")\n >>> ")
	greetPerson("Fred")

	//* Write a function that returns any message, and call it from within
	//  fmt.Println()

	fmt.Println()
	fmt.Println("//* Write a function that returns any message, and call it from within")
	fmt.Println("//  fmt.Println()")
	fmt.Println()
	fmt.Println("Calling fmt.Println(aMessage())\n >>> ")
	fmt.Println(aMessage())
	fmt.Println()


	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer

	fmt.Println("//* Write a function to add 3 numbers together, supplied as arguments, and")
	fmt.Println("//  return the answer")
	fmt.Println()
	fmt.Println("Calling addThreeNumbers(...)\n >>> ")
	a,b,c := 1,2,3
	fmt.Println("a:",a,", b:",b,", c:",c,", sum:", addThreeNumbers(a,b,c) )
	fmt.Println()

	//* Write a function that returns any number
	//* Write a function that returns any two numbers

	fmt.Println()
	fmt.Println("//* Write a function that returns any number")
	fmt.Println()
	fmt.Println("calling function anyNumber()\n >>>")
	fmt.Println("function anyNumber(): ", anyNumber())
	fmt.Println()
	fmt.Println("//* Write a function that returns any two numbers")
	fmt.Println()
	ret1,ret2 := anyTwoNumbers(-1,-2)
	fmt.Println("return from function anyTwoNumbers(int,int): ", ret1, ret2)
	fmt.Println()


	fmt.Println()
	fmt.Println("//* Add three numbers together using any combination of the existing functions.")
	fmt.Println("//  * Print the result")
	fmt.Println(" >>> ")
	fmt.Println("result of adding three numbers:", addThreeNumbers(anyNumber(),1,2) )

	fmt.Println()
	fmt.Println("//* Call every function at least once")
	fmt.Println()

}
