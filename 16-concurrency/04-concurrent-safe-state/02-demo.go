/*
modify the below to execute the increment function concurrently
*/
package main

import (
	"fmt"
	"sync"
)

/* Custom type encapsulating the logic for manipulating the state in a concurrent safe manner */
type SafeCounter struct {
	sync.Mutex
	count int
}

func (sc *SafeCounter) Add(delta int) {
	sc.Lock()
	{
		sc.count += delta
	}
	sc.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
