//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	x, y int
}

func area(rect Rectangle) int {
	return rect.x * rect.y
}

func perimeter(rect Rectangle) int {
	return 2 * (rect.x + rect.y)
}

// type Coordinate struct {
// 	x, y int
// }
// type Rectangle struct {
// 	a Coordinate // top left
// 	b Coordinate // bottom right
// }
// func XsideLen(rect Rectangle) int {
// 	return rect.b.x-rect.a.x
// }
// func YsideLen(rect Rectangle) int {
// 	return rect.a.y-rect.b.y
// }

// func area(rect Rectangle) int {
// 	return XsideLen(rect) * YsideLen(rect)
// }

// func perimeter(rect Rectangle) int {
// 	return 2 * (XsideLen(rect) + YsideLen(rect))
// }

func main() {
	myRect := Rectangle{x: 1, y: 2}
	fmt.Println()
	fmt.Println(" Rectangle: ", myRect)
	fmt.Println("      area: ", area(myRect))
	fmt.Println(" perimiter: ", perimeter(myRect))
	fmt.Println()
	myRect.x = myRect.x * 2
	myRect.y = myRect.y * 2
	fmt.Println("double the rectangle size")
	fmt.Println()
	fmt.Println(" Rectangle: ", myRect)
	fmt.Println("      area: ", area(myRect))
	fmt.Println(" perimiter: ", perimeter(myRect))
	fmt.Println()
}
