package main
import "fmt"

func plus(a int, b int) int{
    return a + b
}

func plusPLus(a, b, c int) int {
    return a + b + c
}

func main(){
    res := plus(1,2)
    fmt.Println("1+2 =", res)
    
    res = plusPLus(1,2,3)
    fmt.Println("1 + 2 + 3 =", res)
}
	
// $ go run functions.go 
// 1+2 = 3
// 1+2+3 = 6


Language Version:  1.10.2