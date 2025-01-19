package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
package main

import "fmt"

func main() {
	a := 5
	for i := a; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= a-i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

