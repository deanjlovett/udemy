package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	
	out := make(chan int, 3) // buffer up to 3 random numbers

	go func() {
		for {
			out <- rand.Intn(max - min + 1) + min
		}
	}()

	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	
	out := make(chan int, 1) // buffer up to 1 random numbers

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max - min + 1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	// seed the RNG so we get new numbers each time
	rand.New( rand.NewSource(time.Now().UnixNano()) )

	randInt := generateRandInt(1, 100)

	fmt.Println()
	fmt.Println("generateRandInt infinite")
	// Get some random integers
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println()
	fmt.Println("generateRandIntn using range")
	randIntnRange := generateRandIntn(4, 1, 10)
	// range loop will automatically break once the channel
	// closes
	for i := range randIntnRange {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("generateRandIntn checked")
	randIntn := generateRandIntn(4, 1, 10)
	for {
		n, isItOpen := <-randIntn
		if !isItOpen {
			break
		}
		fmt.Println(n)
	}
}
