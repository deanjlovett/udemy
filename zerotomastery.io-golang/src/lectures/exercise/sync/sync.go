//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"fmt"

	"os"
	"bufio"
	"strings"
	"sync"
	"unicode"

)

func numberOfLetters(w string) int {
	n := 0
	for _, aRune := range w {
		if unicode.IsLetter(aRune){
			n++
		}
	}
	return n
}

type LetterCount struct {
	count int
	sync.Mutex
}
func (lc *LetterCount)get()int{
	lc.Lock()
	defer lc.Unlock()
	return lc.count
}
func (lc *LetterCount)add(n int){
	lc.Lock()
	defer lc.Unlock()
	lc.count += n
}


func main() {

	var wg sync.WaitGroup
	myscanner := bufio.NewScanner(os.Stdin)

	letterCount := LetterCount{}

	for {
		if myscanner.Scan() {
			aLine := myscanner.Text()
			words := strings.Split(aLine, " ")
			for _, word := range words {
				myword := word
				wg.Add(1)
				//* Input analysis must occur per-word, and each word must be analyzed
				//  within a goroutine
				go func() {
					defer wg.Done()
					letterCount.add(numberOfLetters(myword))
				}()
			} 
		} else {
			if myscanner.Err() == nil {
				break 
			}
		}
	}
	wg.Wait()
	fmt.Println("Total number of letters:", letterCount.get())
}
