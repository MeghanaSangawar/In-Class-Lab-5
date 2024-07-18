package main

import "fmt"

// IsDivisibleByTwo checks if a number is divisible by 2
func IsDivisibleByTwo(n int) bool {
    return n%2 == 0
}

func main() {
    num := 4
    if IsDivisibleByTwo(num) {
        fmt.Printf("%d is divisible by 2\n", num)
    } else {
        fmt.Printf("%d is not divisible by 2\n", num)
    }
}
