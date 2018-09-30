package main
import "fmt"

func main(){
    nums := []int{2,3,4}
    sum := 0
    
    for _, num := range nums {
        sum += num //menambahkan terus pada nilai yang dibungkus oleh array slice di variabel nums declaration
    }
    
    fmt.Println("Sum:",sum)
    
    for i, num := range nums{
        if num == 3 {
            fmt.Println("Index:",i)
        }
    }
    
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v) // %s adalah mencetak key dari map string dan %s\n mencetak nilai dari key yang ada
    }
    for k := range kvs {
        fmt.Println("key:", k)
    }
     for i, c := range "go" {
        fmt.Println(i, c)
    }
}

// $ go run range.go
// sum: 9
// index: 1
// a -> apple
// b -> banana
// key: a
// key: b
// 0 103
// 1 111
