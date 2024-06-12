package main

import (
	"fmt"
	// "math/rand"
	"sync" // a package that allows the usage of wait
	"time"
)

var m = sync.Mutex{}      // this is used in order to allow go routines, in this case, the appending of data into the results slice to occur without any data loss or corruption
var rw = sync.RWMutex{}   // also has a .Rlock and a .RUnlock method on top of the methods that the normal Mutex has
var wg = sync.WaitGroup{} // in short, this is like a counter
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i) // go keyword is used to initiate concurrency
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int) {
	// simulating DB call delay

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Println("The result from the database is", dbData[i])
	m.Lock() // one of the two sync.Mutex methods. Where you place the lock is important as placing it say after delay can remove the idea of concurrency as the dbCall wouldn't continue to occur wasting time
	results = append(results, dbData[i])
	m.Unlock() // the second of two sync.Mutex methods
	wg.Done()
}

// concurrency != parallel execution
