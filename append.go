package main

import "fmt"

func main() {

    var N int
    fmt.Print("Enter the size of the array N: ")
    fmt.Scan(&N)
    array := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Printf("Enter the value for element %d: ", i+1)
        fmt.Scan(&array[i])
    }
    var M int
    fmt.Print("Enter the integer M: ")
    fmt.Scan(&M)
    for i := 0; i < M; i++ {
        temp := array[0]
        for j := 0; j < N-1; j++ {
            array[j] = array[j+1]
        }
        array[N-1] = temp
    }

    fmt.Println("Array after moving the first", M, "elements to the end:")
    for _, value := range array {
        fmt.Print(value, " ")
    }
    fmt.Println()
}