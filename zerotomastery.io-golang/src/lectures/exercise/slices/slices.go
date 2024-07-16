//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

// import "fmt"
import (
	"fmt"
	"os"
	"strings"
)

type Part string

//
//* Create a function to print out the contents of the assembly line
//
func printAssemblyLine(assemblyLine []Part, longFmt bool) {
	if(!longFmt) {
		fmt.Println("value:", assemblyLine)
	} else {
		fmt.Println("value:")
		for idx, elem := range assemblyLine {
			fmt.Println("\t",idx,elem)
		}
	}
	fmt.Println()
}

func main() {
	longFmt := false
	args := os.Args[1:]
	if len(args)>0 {
		flag := strings.ToLower(args[0])
		switch flag {
		case "-l","--long":
			longFmt = true
			fmt.Println("using long formt")
		case "-s","--short":
			longFmt = false
			fmt.Println("using short formt")
		default:
			fmt.Printf("unknow flag:[%s], using short formt",args[0])
		}
	}

	fmt.Println()
	fmt.Println("assembly line:")
	fmt.Println()

	//  - Print out the contents of the assembly line at each step

	//  - Create an assembly line having any three parts
	
	assemblyLine := []Part{"one", "two", "three"}
	fmt.Println("  cmd:", "assemblyLine := []Part{\"one\",\"two\",\"three\"}")
	printAssemblyLine(assemblyLine,longFmt)

	//  - Add two new parts to the line

	assemblyLine = append(assemblyLine, "four", "five")
	fmt.Println("  cmd:", "assemblyLine = append(assemblyLine,\"four\",\"five\")")
	printAssemblyLine(assemblyLine,longFmt)

	//  - Slice the assembly line so it contains only the two new parts

	assemblyLine = assemblyLine[3:]
	fmt.Println("  cmd:", "assemblyLine = assemblyLine[3:]")
	printAssemblyLine(assemblyLine,longFmt)
}
