package main

import "fmt"

func main(){
    if 7%2 == 0 {
        fmt.Println("7 is Even")
    } else {
        fmt.Println("7 is Odd")
    }
    
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
    
    if num := 9; num < 0 {
        fmt.Println(num, "Is Negative")
    } else {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num,"has multiple digits")
    }
}



// $ go run if-else.go 
// 7 is odd
// 8 is divisible by 4
// 9 has 1 digit

