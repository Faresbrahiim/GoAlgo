package main

import "fmt"

func main() {
	a := 5
	for i := 1; i <= a; i++ {
		for k := 1; k < i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*(a-i)+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

*********
 *******
  *****
   ***
    *