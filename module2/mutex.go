package main

import (
	"fmt"
	"sync"
	"time"
)

func oneTime() {
	fmt.Println("hi")
}

func mutexMain() {

	var wt sync.WaitGroup
	var mt sync.Mutex
	var oe sync.Once
	var rw sync.RWMutex
	counter := 0

	for i := 1; i <= 1; i++ {
		wt.Add(1)

		go func() {
			defer wt.Done()
			mt.Lock()
			oe.Do(oneTime)
			counter++
			mt.Unlock()
		}()

	}
	wt.Wait()
	fmt.Println(counter)

	mun := 0

	go func() {

		rw.RLock()
		defer rw.RUnlock()
		tar := mun
		fmt.Println(tar)

	}()

	go func() {

		rw.Lock()
		mun++
		defer rw.Unlock()

	}()

	time.Sleep(3 * time.Second)

}
