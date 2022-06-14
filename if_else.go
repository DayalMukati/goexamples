package main

import "fmt"

func main() {

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 {
        fmt.Println("8 is even")
    } else {
        fmt.Println("8 is odd")
    }

    if num := 14; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}