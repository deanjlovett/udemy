package main

import (
	// "context"
	"bufio"
	"fmt"
	// "time"
	"log"
	"math/rand"
	"os"
	"path/filepath" // filepath.Walk
	"strings"
	"sync"
	"time"
)

func mywait(delay int) {
	time.Sleep(time.Duration(rand.Intn(delay)) * time.Millisecond)
}

func optionParDir(folder, match string, parentwg *sync.WaitGroup) {
	defer parentwg.Done()
	var wg sync.WaitGroup
	// fmt.Printf("searching folder: %v ... start ...\n",folder)

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		filename := folder + "/" + f.Name()
		wg.Add(1)
		if f.IsDir() {
			go optionParDir(filename,match,&wg)
		} else {
			go optionOneFile(filename,match,&wg)
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
	}
	for _, f := range files {
		filename := folder + "/" + f.Name()
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
				"Line#:%v, File:%v, Matching line:%v\n\n",
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

func mainOptionTwo(searchStr, searchDir string) {
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("Option 2: filepath.Walk")
	fmt.Println()


	var wg sync.WaitGroup

	searchFile := func( path string, match string, wg *sync.WaitGroup) <-chan string{
		defer wg.Done()
		out := make(chan string)
		out <- fmt.Sprintf("looking for {%v} in file {%v}\n",match,path)
		mywait(100)
		return out
	}
	refunc := func( path string, info os.FileInfo, err error ) error {
		if err != nil {
			return err
		}
		isDirStr := func(isDir bool) string {
			if isDir { return "D" } else { return "_"}
		}
		fmt.Println(isDirStr(info.IsDir()), path, info.Size())
		go searchFile(path,searchStr,&wg)
		return nil
	}

	err := filepath.Walk( searchDir, refunc )
	// 	func(
	// 		path string, 
	// 		info os.FileInfo, 
	// 		err error,
	// 	) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(path, info.Size())
	// 		return nil
	// 	},
	// )
	if err != nil {
		fmt.Println(err)
	}
}

func mainOptionThree(searchStr, searchDir string){
	
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("Option 3: os.File.Readdir")
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
	fmt.Println()
	fmt.Println("===========================") 
	fmt.Println("from...")
	fmt.Println("https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html")
	fmt.Println()

	
	
	mainOptionOne(searchStr, searchDir)

	// mainOptionTwo(searchStr, searchDir)

	// mainOptionThree(searchStr, searchDir)


}