// go routine is a specila type of function it can be executed independently
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// fmt.Println("index", i)
		wg.Add(1) // mention that only one go routine
		go worker(i, &wg)
	}
	wg.Wait()
}
