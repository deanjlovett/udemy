//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type product struct {
	name string
	price int
}

func printStats(prods [4] product) {
	var totalCost, totalProds int

	for i:=0;i<len(prods);i++{
		var prod = prods[i]
		if(prod.name=="") { continue }
		totalCost += prod.price
		totalProds++
	}
	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items

	fmt.Println();
	fmt.Println("  The last item on the list:", prods[totalProds-1])
	fmt.Println("  The total number of items:", totalProds)
	fmt.Println("The total cost of the items:", totalCost)
	fmt.Println();

}

func main() {
	prodArray := [4] product{
		{"One",1}, 
		{"Two",2}, 
		{"Three",3},
	}
	printStats(prodArray);
	prodArray[3]=product{"Four",4}
	printStats(prodArray);
}
