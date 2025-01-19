package main

import "fmt"

func main() {
	a := 5
	for i := a; i >= 1; i-- { // Outer loop: Rows in reverse order
		for j := 1; j <= a-i; j++ { // Spaces increase with each row
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ { // Stars decrease with each row
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


