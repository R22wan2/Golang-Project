package main

// import (
// 	"fmt"
// )

// !-0
// func sleeper(numb int, show string) {

// 	for i := 0; i < numb; i++ {
// 		time.Sleep(300 * time.Millisecond)
// 		fmt.Printf("%s  its value is %d\n", show, i)
// 	}

// }
// !-1
// func add(value chan int) {

// 	time.Sleep(5 * time.Second)
// 	value <- 90

// }
// !-2
// ! PIPE
// func value_geter(width float64, length float64, square chan float64) {

// 	square <- width * length

// }

// func show_output(decision float64) {
// 	fmt.Printf("square is %.2f.", decision)
// }

func main() {
	// !-0
	// fmt.Println("it begins")
	// go sleeper(5, "bottom")
	// go sleeper(7, "right")
	// go sleeper(9, "left")
	// go sleeper(8, "top")
	// !-1
	// channel := make(chan int)

	// go add(channel)

	// total := <-channel

	// fmt.Println(total * total)
	// !-2

	// var length float64
	// var width float64
	// fmt.Print("please enter length : ")
	// fmt.Scan(&length)
	// fmt.Print("please enter width : ")
	// fmt.Scan(&width)

	// square := make(chan float64)

	// go value_geter(width, length, square)

	// final_value := <-square

	// show_output(final_value)

	// ! PIPE

	// inp := make(chan int)
	// out := make(chan int)

	// done := make(chan struct{})

	// go numberGenerateor(inp)
	// go addValue(inp, out)
	// go printNumber(out, done)

	// <-done
	// 	first_val := make(chan int, 3)
	// 	first_val <- 29
	// 	first_val <- 2
	// 	first_val <- 20
	// 	fmt.Printf("the value %d\n", <-first_val)
	// 	fmt.Printf("the value %d\n", <-first_val)
	// 	fmt.Printf("the value %d\n", <-first_val)

	// }

	// func numberGenerateor(numb1 chan int) {

	// 	for i := 1; i <= 5; i++ {
	// 		numb1 <- i
	// 	}
	// 	close(numb1)

	// }

	// func addValue(numb1 chan int, out chan int) {
	// 	for numb := range numb1 {

	// 		if !(numb%2 == 0) {
	// 			out <- numb * numb
	// 		} else {
	// 			out <- numb
	// 		}

	// 	}
	// 	close(out)
	// }

	// func printNumber(final_value chan int, done chan struct{}) {

	// 	for numb := range final_value {
	// 		fmt.Printf("the value %d\n", numb)
	// 	}
	// 	done <- struct{}{}
	// a := 1
	// if a == 1 {
	// 	selecter()
	// } else {
	// 	mutexMain()
	// }
	filer()

}
