package main

import "fmt"

func printMap(aMap map[string]int){
	fmt.Println()
	for idx,elem := range aMap {
		fmt.Println(idx,elem)
	}
	fmt.Println()
}

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	// printMap(shoppingList)
	fmt.Println(shoppingList)
	// for idx,elem := range shoppingList {
	// 	fmt.Println(idx,elem)
	// }
	shoppingList["eggs"] += 1

	// printMap(shoppingList)
	fmt.Println(shoppingList)

	delete(shoppingList,"milk")
	fmt.Println(shoppingList)

	fmt.Println("Milk deleted, new list:", shoppingList)

	cereal, found := shoppingList["cereal"]
	fmt.Print("Need cereal?  ")
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yes", cereal, "boxes")
	}
	totalItems := 0
	for _, amount:= range shoppingList{
		totalItems += amount
	}
	fmt.Println("There are", len(shoppingList), "entries on the shopping list.")
	fmt.Println("There are", totalItems, "on the shopping list.")

}
