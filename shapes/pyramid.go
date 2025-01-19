package main

import "fmt"

func main() {
	n := 7
	mid := n / 2

	for i := 1; i <= mid+1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := mid; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
