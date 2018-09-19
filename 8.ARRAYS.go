package main

import "fmt"

func main(){
    
    var a [5]int // variabel a mempunyai panjang array 5 buah yang memiliki tipe data integer
    fmt.Println("emp:", a)
    
    a[4] = 100
    fmt.Println("Set:", a) //kita akan set nilai angka array pada urutan ke 4 menjadi 100 dengan tipe integer
    fmt.Println("get:", a[4]) // kita get atau panggil atau memunculkan value array pada var a dengan urutan ke 4
    
    fmt.Println("Len:", len(a)) // melihat panjang array pada variabel a integer
    
    b:= [5]int{1,2,3,4,5} // menset seluruh urutan array yang memiliki array dengan nilai array 1,2,3,4,5
    fmt.Println("dcl:",b) 
    
    var twoD [2][3]int // membuat array 2 dimensi
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

// $ go run arrays.go
// emp: [0 0 0 0 0]
// set: [0 0 0 0 100]
// get: 100
// len: 5
// dcl: [1 2 3 4 5]
// 2d:  [[0 1 2] [1 2 3]]


Language Version:  1.10.2