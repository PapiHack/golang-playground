package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition - LearnCodeOnline.in")

	waitGroup := &sync.WaitGroup{}

	mutex := &sync.RWMutex{}

	var score = []int{0}

	waitGroup.Add(3) // We can pass directly the total number of WG
	go func (wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		mutex.Unlock()
		waitGroup.Done()
	}(waitGroup, mutex)

	// waitGroup.Add(1)
	go func (wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		mutex.Unlock()
		waitGroup.Done()
	}(waitGroup, mutex)

	go func (wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		waitGroup.Done()
	}(waitGroup, mutex)

	go func (wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("SCORE")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		waitGroup.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	fmt.Println(score)
}
