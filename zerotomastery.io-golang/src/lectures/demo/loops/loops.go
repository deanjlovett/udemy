package main

import "fmt"

func main() {
	sum := 0
	fmt.Println()
	fmt.Println("Begin")
	fmt.Println("=====")
	fmt.Println("Sum is", sum)
	for i:=1;i<=10; i++ {
		sum += i
		fmt.Println("increment. Sum is", sum)
	}
	fmt.Println("=====")
	for sum > 10 {
		sum -= 5
		fmt.Println("decrement. Sum is", sum)
	}
	fmt.Println("=====")
	fmt.Println("  End")
	fmt.Println()
}
