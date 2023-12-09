package main

import (
	"fmt"
	//"internal/itoa"
	"math/rand"
	"sync"
	"time"

	//"golang.org/x/tools/go/analysis/passes/defers"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex // embedded directly in struct. can be accessed directly
}
// todo: add mutex protected access fuctions
// getHits or Hits
// incHits
//
func getHits(hits *Hits)int{
	hits.Lock()
	defer hits.Unlock()
	return hits.count
}
func incHits(hits *Hits){
	hits.Lock()
	defer hits.Unlock()
	hits.count += 1
}

func serve(wg *sync.WaitGroup, hits *Hits, interation int){
	wait()
	defer wg.Done()
	// original
	//
	// hits.Lock()
	// defer hits.Unlock()
	// hits.count += 1

	// update
	incHits(hits)

	fmt.Println("served intertion", interation)
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	// myrand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	hitCounter := Hits{}
	for i := 0; i<20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}
	fmt.Printf("waiting for goroutines...\n\n")
	wg.Wait()

	// original
	//
	// hitCounter.Lock()
	// totalHits := hitCounter.count
	// hitCounter.Unlock()

	// update
	totalHits := getHits(&hitCounter)

	fmt.Printf("\ntotal hits = %d\n",totalHits)
}
