package main

import (
	"fmt"
	"sync"
)

func runTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Running task %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runTask(i, &wg)
	}

	wg.Wait()
}