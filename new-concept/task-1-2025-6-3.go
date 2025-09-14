package main

import (
	"fmt"
)

// generic

type boxType[T any] struct {
	value T
}

func (result boxType[T]) Get_Value(p_value T) string {
	return fmt.Sprint(result.value, "---hello---", p_value)
}

// main
func Generic_Main() {

	hel := boxType[int]{value: 20}

	fmt.Print(hel.Get_Value(87))

}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// odd or even

func Odd_Or_Even(val int) {

	if val <= 0 {
		fmt.Println("you can't send 0 or under 0")
		return
	}

	result := val % 2

	if result == 0 {
		fmt.Println("givenffff value is even")
	} else {
		fmt.Println("given value is odd")
	}

}
