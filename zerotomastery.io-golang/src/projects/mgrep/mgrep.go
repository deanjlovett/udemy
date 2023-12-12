package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// func mywait(delay int) {
// 	time.Sleep(time.Duration(rand.Intn(delay)) * time.Millisecond)
// }

func optionParDir(folder, match string, parentwg *sync.WaitGroup) {
	defer parentwg.Done()
	var wg sync.WaitGroup
	// fmt.Printf("searching folder: %v ... start ...\n",folder)

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range files {
			// filename := folder + "/" + f.Name()
			filename := folder + string(os.PathSeparator) + f.Name() 

			wg.Add(1)
			if f.IsDir() {
				go optionParDir(filename,match,&wg)
			} else {
				go optionOneFile(filename,match,&wg)
			}
		}
	}
	wg.Wait()
	// fmt.Printf("searching folder: %v ... stop ...\n",folder)

}

func optionOneDir(folder, match string, wg *sync.WaitGroup) {
	defer wg.Done()

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		// filename := folder + "/" + f.Name() // << works for Mac OS & Linux
		// filename := fmt.Sprintf("%v%v%v",folder,os.PathSeparator,f.Name())
		filename := folder + string(os.PathSeparator) + f.Name() 
		wg.Add(1)
		if f.IsDir() {
			go optionOneDir(filename,match,wg)
		} else {
			go optionOneFile(filename,match,wg)
		}
	}
}
func optionOneFile(filename, match string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
		return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
	linenumber := 0
	for scanner.Scan() {
		linenumber++
		line := scanner.Text()
		if strings.Contains(line,match) {
			fmt.Printf(
				// "Line#: %v, File: %v, Matching line: %v\n\n",
				"Line#: %v\nFile: %v\nMatching line: %v\n\n",
				linenumber,
				filename,
				line,
			)
		}
    }
	// mywait(100)
	// fmt.Printf("searching file: %v for %v complete !\n",filename,match)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


func mainOptionOne(searchStr, searchDir string){
	var wg sync.WaitGroup
	wg.Add(1)
	go optionParDir(searchDir, searchStr, &wg)
	wg.Wait()
	fmt.Println()
}

func main(){

	if len(os.Args)<3 {
		return
	}
	searchStr := os.Args[1]
	searchDir := os.Args[2]
	fmt.Println()
	fmt.Println("searchStr:", searchStr)
	fmt.Println("searchDir:", searchDir)
	fmt.Println()
	fmt.Println("===========================") 
	fmt.Println()
	
	mainOptionOne(searchStr, searchDir)

}