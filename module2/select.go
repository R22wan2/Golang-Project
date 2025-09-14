package main

import (
	"fmt"
	"time"
)

func genric[T any](book T) T {
	var man T = book
	fmt.Print(man)
	return book

}

func selecter() {

	fmt.Println(genric(9))

	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {

		time.Sleep(4 * time.Second)
		channel1 <- 5
		close(channel1)

	}()
	go func() {

		time.Sleep(3 * time.Second)
		channel2 <- 9
		close(channel2)
	}()

	select {
	case val := <-channel1:
		fmt.Printf("the number is %d", val)

	case val := <-channel2:
		fmt.Printf("the number is %d", val)
	}
	time.Sleep(6 * time.Second)
}
