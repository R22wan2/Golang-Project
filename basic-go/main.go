package main

import "fmt"

type datas_types struct {
	name        string
	age         int
	height      float32
	gender      string
	achievement struct {
		book_no int
	}
}

func (details datas_types) getFulldetails() string {

	decision_of_gender := ""

	if details.gender == "male" {
		decision_of_gender = "mr"
	} else if details.gender == "female" {
		decision_of_gender = "ms"
	} else {
		decision_of_gender = "misr"
	}

	return_value := fmt.Sprintf("%s %s and your age is %d your height %.2f achievements books of %d", decision_of_gender, details.name, details.age, details.height, details.achievement.book_no)

	return return_value

}

// func nameAdd(sum1, sum2 int) int {
// 	return sum1 + sum2
// }

func maisn() {

	//Function
	// fmt.Println(nameAdd(9, 2))

	person1 := datas_types{
		name:   "riswan",
		age:    20,
		height: 177.5,
		gender: "male",
		achievement: struct{ book_no int }{
			book_no: 9,
		},
	}

	person2 := datas_types{
		name:   "rilwan",
		age:    23,
		height: 170.5,
		gender: "male",
		achievement: struct{ book_no int }{
			book_no: 89,
		},
	}

	fmt.Println(person1.getFulldetails())
	fmt.Println(person2.getFulldetails())

}
