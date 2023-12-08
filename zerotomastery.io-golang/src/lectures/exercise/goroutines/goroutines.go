//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"

	// "io"
	"os"
	"strconv"
	"sync"
	"time"
	// "math/rand"
)

func panicCheck(err error) {
	if err != nil {
		fmt.Println("Error:",err)
		panic(err)
	}
}
func check(err error) {
	if err != nil {
		fmt.Println("Error:",err)
	}
}
func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	var wg sync.WaitGroup
	

	// cin := make(chan int,2*(len(files)+1))
	cin := make(chan int, (len(files) + 1))
	// make buffer large enough for all the file readers

	fn := func(tag int, fname string, ci <-chan int, wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		sum := 0
		f, err := os.Open(fname)
		panicCheck(err)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			sline := scanner.Text()
			i, err := strconv.Atoi(sline)
			if err == nil {
				sum += i
			}else{
				fmt.Printf("Error file: %v, strlen: %v, string: \"%v\", msg{%v}\n", fname,len(sline),sline, err)
			}
		}
		// added delay to test waiting for threads to start
		//
		// time.Sleep(time.Duration(rand.Intn(1000)+1) * time.Millisecond)

		fmt.Println("File:",fname,"sum =", sum)
		// cin<-tag
		cin <- sum
	}

	for idx, fileName := range files {
		fmt.Println("  inside for loop: idx:", idx, "fileName:", fileName)
		go fn(idx+1, fileName, cin, &wg)
	}

	ncin := len(cin)
	var waitCount time.Duration = 1
	for ncin == 0 {
		fmt.Println("  waiting for threads to write to channel. wait time:", waitCount)
		time.Sleep(waitCount * time.Millisecond)
		waitCount *= 2
		ncin = len(cin)
	}
	wg.Wait()
	sum := 0
	for ncin > 0 {
		// tag := <-cin
		sum += <-cin
		// fmt.Println("  inside for loop: reading from channel. tag:",tag)
		fmt.Println("  inside for loop: reading from channel. # in buffer:", ncin)
		ncin = len(cin)
	}
	close(cin)
	fmt.Println("total sum =", sum)
}
