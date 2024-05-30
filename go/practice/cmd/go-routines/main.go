package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// Add 1 to the wait group
		wg.Add(1)
		go dbCall(i)
	}
	// Wait for the wait group to reach 0 (currently 5 as the loop has ran 5 times)
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results)
}

func dbCall(i int) {
	// Simulate DB Call
	var delay float32 = 0.000000001
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
