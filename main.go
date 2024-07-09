package main

import "fmt"

func main() {
	// var firstName string
	// var lastName string
	// var email string
	// arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
	// 	{"C++", "Go", "HTML"}}

	arr := []int{1, 2, 3, 4, 5}

	arrLastIndex := len(arr) - 1

	slices := arr[:arrLastIndex+1]

	fmt.Println("slices =>", slices)
	fmt.Println("arr =>", arr)
	// Accessing the values of the
	// array Using for loop
	// fmt.Println("Elements of Array 1")
	// for x := 0; x < 3; x++ {
	// 	fmt.Println(arr[x])
	// 	for y := 0; y < 3; y++ {

	// 		fmt.Println(arr[x][y])
	// 	}
	// }

	// fmt.Println("Enter your first name: ")
	// fmt.Scan(&firstName)

	// fmt.Println("Enter your last name: ")
	// fmt.Scan(&lastName)

	// fmt.Println("Enter your email: ")
	// fmt.Scan(&email)

	// fmt.Printf("firstName is %v,lastName is %v and the email is %v ", firstName, lastName, email)

}
