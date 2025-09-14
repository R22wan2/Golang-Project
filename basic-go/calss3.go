package main

import "fmt"

// map

func add(sum1 int, sum2 int) int {
	return sum1 + sum2
}
func sub(sum1 int, sum2 int) int {
	return sum1 - sum2
}

func combaine_access(value1, value2 int, decision_of_operation func(int, int) int) string {

	answer := decision_of_operation(value1, value2)

	return fmt.Sprintf("answer is %d", answer)

}

func plus() func() int {
	a := 0
	return func() int {
		a += 2
		return a
	}
}

func nana() {

	// options := make(map[string]int)
	// options["name"] = 100
	// fmt.Println(options["name"])
	// age := map[string]int{
	// 	"no": 9,
	// }
	// fmt.Println(age)
	// nested map
	// !nested map
	// marks := map[string]map[string]string{
	// 	"age": {
	// 		"manfd": "null",
	// 	},
	// 	"man": {
	// 		"ms": "hs",
	// 	},
	// }

	// fmt.Println(marks["age"]["manfd"])
	// ! high order function

	// this defer

	defer fmt.Println("_____________________")

	fmt.Println(combaine_access(627, 99, add))
	fmt.Println(combaine_access(627, 99, sub))
	var function_access func() int = plus()
	fmt.Println(function_access())
	fmt.Println(function_access())
	fmt.Println(function_access())
	fmt.Println(function_access())
	fmt.Println(function_access())

}

// !Pointer

func change_pointer_value(valuer *int) {

	*valuer = 2

}

type avarage struct {
	name string
}

func (changer_of_name *avarage) change_name() {
	changer_of_name.name = "sds"
}

func pointer() {

	mon := 9
	change_pointer_value(&mon)

	fmt.Println(mon)

	humble := avarage{
		name: "rilwan",
	}
	fmt.Println(humble.name)
	humble.change_name()
	fmt.Println(humble.name)

}
