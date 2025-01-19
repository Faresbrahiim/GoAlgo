package main

import "fmt"

func main() {
    // Step 1: Ask for the size of the array N
    var N int
    fmt.Print("Enter the size of the array N: ")
    fmt.Scan(&N)

    // Step 2: Create an array of size N
    array := make([]int, N)

    // Step 3: Fill the array with user input
    for i := 0; i < N; i++ {
        fmt.Printf("Enter the value for element %d: ", i+1)
        fmt.Scan(&array[i])
    }

    // Step 4: Ask for the integer M
    var M int
    fmt.Print("Enter the integer M: ")
    fmt.Scan(&M)

    // Step 5: Move the first M elements to the end of the array
    // Without using append or slicing
    for i := 0; i < M; i++ {
        // Store the first element temporarily
        temp := array[0]

        // Shift all elements to the left by one position
        for j := 0; j < N-1; j++ {
            array[j] = array[j+1]
        }

        // Place the stored element at the end
        array[N-1] = temp
    }

    // Step 6: Print the modified array
    fmt.Println("Array after moving the first", M, "elements to the end:")
    for _, value := range array {
        fmt.Print(value, " ")
    }
    fmt.Println()
}