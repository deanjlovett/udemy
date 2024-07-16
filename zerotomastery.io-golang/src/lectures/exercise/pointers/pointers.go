//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active = true
	Inactive = false
)
type SecurityTag bool

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
type SecTag struct {
	name string
	tag SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(item *SecTag){
	item.tag = Active // true
}
func deactivate(item *SecTag){
	item.tag = Inactive //false
}
//* Create a checkout() function which can deactivate all tags in a slice
// func checkout(items *[]SecTag){
// 	for i:=0;i<len(*items);i++{
// 		(*items)[i].tag = false
// 	}
// }
func checkout(items []SecTag){
	for i:=0;i<len(items);i++{
		deactivate(& items[i]) //items[i].tag = false
	}
}
func checkin(items []SecTag){
	for i:=0;i<len(items);i++{
		activate(& items[i]) //items[i].state = true
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	//  - Deactivate any one security tag in the array/slice
	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change

	items := []SecTag{
		{"one",true},
		{"two",true},
		{"three",true},
		{"four",true},
	}
	fmt.Println(items)
	checkin(items)
	fmt.Println(items)

	deactivate(&items[3])
	fmt.Println(items)

	// checkout(&items)
	checkout(items)
	fmt.Println(items)

}
