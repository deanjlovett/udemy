package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
)

// func mywait(delay int) {
// 	time.Sleep(time.Duration(rand.Intn(delay)) * time.Millisecond)
// }

// create a new WaitGroup to be used by child threads / go routines
// wait for child threads to finish before leaving method

func optionParDir(folder, match string, parentwg *sync.WaitGroup) {
	defer parentwg.Done()
	var wg sync.WaitGroup

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range files {
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
}

// new child threads - go routines us WaitGroup passed in
// do not wait in this method for child threads to complete
//
func optionOneDir(folder, match string, wg *sync.WaitGroup) {
	defer wg.Done()

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
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
			fmt.Printf( "%v[%v]:%v\n", filename, linenumber, line )
		}
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


func mainOptionOne(searchStr, searchDir string){

	fmt.Println("===========================") 
	fmt.Println("===========================") 
	fmt.Println()
	fmt.Println("searchStr:", searchStr)
	fmt.Println("searchDir:", searchDir)
	fmt.Println()
	fmt.Println("===========================") 
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(1)
	go optionParDir(searchDir, searchStr, &wg)
	wg.Wait()
	fmt.Println()
}

func main(){

	helpFlags := []string {"-h","-?","--help"}
	usage := "\nusage mgrep search_string [folder ... ]\n"
	checkForHelpFlag := func (searchStr string)bool{
		if slices.Contains(helpFlags,strings.ToLower(searchStr)) {
			return true //os.Exit(0)
		}
		return false
	}

	searchStr := ""
	fmt.Println("len(os.Args):",len(os.Args),os.Args)
	fmt.Println()
	if len(os.Args)<=1 {
		fmt.Println("\n   need to provide at least provide a search string.")
		fmt.Println(usage)
	} else if len(os.Args) == 2 {
		searchStr = os.Args[1]
		if checkForHelpFlag(searchStr) {
			fmt.Println(usage)
			os.Exit(0)
		}
		mainOptionOne(searchStr, ".")
	} else {
		searchStr       = os.Args[1]
		searchArray := os.Args[1:]
		for _,aString := range searchArray {
			if checkForHelpFlag(aString) {
				fmt.Println(usage)
				os.Exit(0)
			}
		}
		searchDirArray := os.Args[2:]
		for _,searchDir := range searchDirArray {
			mainOptionOne(searchStr, searchDir)
		}
	}
}

