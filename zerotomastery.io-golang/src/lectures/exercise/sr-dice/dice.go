//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

//import "time"

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	// randomGenerator := new(rand.NewSource(time.Now().UnixNano()))

	dice, sides := 2, 12
	maxRolls := 1000

	args := os.Args[1:]
	if len(args) >= 1 {
		test, e := strconv.Atoi(args[0])
		if e == nil {
			dice = test
		}
	}
	if len(args) >= 2 {
		test, e := strconv.Atoi(args[1])
		if e == nil {
			sides = test
		}
	}
	if len(args) >= 3 {
		test, e := strconv.Atoi(args[2])
		if e == nil {
			maxRolls = test
		}
	}

	// rolls := 1
	specialSnakeEyes := 0
	specialLucky7 := 0
	specialOdds := 0
	specialEvens := 0

	fmt.Println()
	fmt.Println("     # of  dice:", dice)
	fmt.Println("     # of sides:", sides)
	fmt.Println(" Max # of rolls:", maxRolls)
	fmt.Println()
	// for r := 1; r <= rolls; r++ {
	r := 1
	for ; (specialSnakeEyes == 0 || specialLucky7 == 0) && r <= maxRolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", Die #", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes!")
			specialSnakeEyes++
		case sum == 7:
			fmt.Println("Lucky 7")
			specialLucky7++
		case sum%2 == 0:
			fmt.Println("Even")
			specialEvens++
		case sum%2 == 1:
			fmt.Println("Odd")
			specialOdds++
		}
		fmt.Println()

		//  - "Snake eyes": when the total roll is 2, and total dice is 2
		//  - "Lucky 7": when the total roll is 7
		//  - "Even": when the total roll is even
		//  - "Odd": when the total roll is odd

	}
	fmt.Println()
	fmt.Println("cmdline args:", len(args), args)
	fmt.Println()
	fmt.Println("     # of  dice:", dice)
	fmt.Println("     # of sides:", sides)
	fmt.Println(" Max # of rolls:", maxRolls)
	fmt.Println()
	fmt.Println("Total # of Snake Eyes =", specialSnakeEyes)
	fmt.Println("Total # of    Lucky 7 =", specialLucky7)
	fmt.Println()
	fmt.Println("Total # of       Odds =", specialOdds)
	fmt.Println("Total # of      Evens =", specialEvens)
	fmt.Println("          Total rolls =", r)

}
