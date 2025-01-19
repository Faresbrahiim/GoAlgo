package main

import "fmt"

func main() {
	var n, number int
	cP := 0    // Counter for prime numbers
	cPerf := 0  // Counter for perfect numbers

	fmt.Print("How many numbers: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter number: ")
		fmt.Scan(&number)

		// Prime number check
		flag := true
		if number <= 1 {
			flag = false
		} else {
			for j := 2; j <= number/2; j++ {
				if number%j == 0 { // Check divisibility
					flag = false // Not a prime number
					break
				}
			}
		}

		// If number is prime, increment prime counter
		if flag {
			cP++
		}

		// Perfect number check
		sum := 0
		for k := 1; k < number; k++ { // Check divisors excluding the number itself
			if number%k == 0 {
				sum += k
			}
		}
		if sum == number {
			cPerf++ // Perfect number found
		}
	}

	fmt.Printf("Total prime numbers: %d\n", cP)
	fmt.Printf("Total perfect numbers: %d\n", cPerf)
}
