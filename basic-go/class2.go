package main

// import "fmt"

// ^  29.4.2025
//! --------------------------------------------------------------------------INTER FACE
// type myclassical_data interface {
// 	abdDatas(data string) string
// }

// type manager struct {
// 	name string
// }

// func (details manager) abdDatas(data string) string {

// 	return details.name + " : " + data
// }
// func sender(inputs myclassical_data) {
// 	fmt.Print(inputs.abdDatas("leona"))
// }
// ^  30.4.2025
// !------------------------------USE INTERFACE OR ANY TO GIVEBACK  TYPE AND ERROR HANDLING
// func division_number(val1, val2 float64) (float64, error, any) {

// 	if val2 == 0 {
// 		return 0, fmt.Errorf("error you are gave value of %d", 0), 0
// 	}
// 	a := 9
// 	return val1 / val2, nil, a
// }

// ^  1.5.2025

// ! ARRAY
// func showbooks(books []string) {

// 	for _, book := range books {
// 		fmt.Println(book)
// 	}

// }

// * add values

// func addingvalues(counts []int) int {

// 	total := 0

// 	for _, i := range counts {

// 		total += i

// 	}

// 	return total

// }

// func matrixshow(matrixs [][]int) {

// 	for i := 0; i < len(matrixs); i++ {
// 		for j := 0; j < len(matrixs[i]); j++ {
// 			fmt.Print(matrixs[i][j]+9, "\t")
// 		}
// 		fmt.Print("\n\n")
// 	}

// }

func main() {
	// ^  29.4.2025
	//! -------------------INTERFACE
	// inputs := manager{
	// 	name: "stringer",
	// }
	// sender(inputs)

	// ^  30.4.2025
	// !--------------------USE INTERFACE OR ANY TO GIVEBACK  TYPE AND ERROR HANDLING
	// val1, val2 := 5.0, 2.0

	// final_result, err, n := division_number(val1, val2)
	// _, ok := n.(int)

	// fmt.Println(ok)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("answer : %.2f", final_result)

	// ^  1.5.2025
	// ! ARRAY
	// books := []string{
	// 	"rich dad and poor dad",
	// 	"psychology of money",
	// }

	// showbooks(books)

	// phones := make([]int, 0, 100)

	// for inser_value := 1; inser_value <= 100; inser_value++ {

	// 	phones = append(phones, inser_value)

	// }

	// final_value := addingvalues(phones)

	// fmt.Printf(" the total value is %d.", final_value)

	// ! Matrix

	// matrix_value := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{1, 2, 3, 4, 5},
	// }

	// matrixshow(matrix_value)

	// !Map and aonther like high order function nested map defer and some function type

	// nana()
	pointer()

}
