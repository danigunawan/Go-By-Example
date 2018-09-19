package main

import "fmt"
import "time"

func main(){
    i := 2
    
    switch i {
        case 1:
        fmt.Println("One")
        case 2:
        fmt.Println("Two")
        case 3:
        fmt.Println("Three")
    }
    
    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
        default:
        fmt.Println("Its a Weekday")
    }
    
    t := time.Now()
    
    switch {
        case t.Hour() < 12:
        fmt.Println("Its before noon")
        default:
        fmt.Println("Its after noon")
    }
    
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

// $ go run switch.go 
// Write 2 as two
// It's a weekday
// It's after noon
// I'm a bool
// I'm an int
// Don't know type string