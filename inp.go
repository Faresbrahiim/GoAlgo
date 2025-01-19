package main

import "fmt"

func main() {
	// Declare all variables locally within the main function
	x := 10
	y := 3
	z := 7
	result := 0

	// Declare `i` and `j` outside the loops
	var i, j int

	// Outer loop using `i`
	for i = 0; i < x; i++ {
		fmt.Print("Outer Loop: ", i, " | ")
		// Inner loop using `j`
		for j = y; j > 0; j-- {
			fmt.Print("Inner Loop: ", j, " | ")
			if j%2 == 0 {
				z += 2
			} else {
				z -= 1
			}
			result += z
		}
		fmt.Println("Result so far:", result)
	}

	fmt.Println("Final Result:", result)
	fmt.Println("Final z:", z)
	fmt.Println("Final i:", i)
	fmt.Println("Final j:", j)
}
