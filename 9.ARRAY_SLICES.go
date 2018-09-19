// ARRAY SLICES
// Slices adalah tipe data utama di Go, intinamah urutan array yang lebih kuat.
// mirip array numeric, ngga miripp array associatif ngga juga array multidimensional
package main

import "fmt"

func main(){
    s := make([]string, 3) // membuat bungkus array berjumlah 3 urutan
    fmt.Println("Emp", s)
    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("Set:", s)
    fmt.Println("Get:", s[2])
    
    fmt.Println("Len:", len(s))
    
    s = append(s, "d") //menyisipkan value "d" di array berikutnya setelah array s[2] dengan nilai "c"
    s = append(s, "e", "f") //menyisipkan value "e" dan "f" di array berikutnya setelah array yang memiliki value "d"
    fmt.Println("apd:", s) // echo dan passing lah
    
    c := make([]string, len(s)) // hitung panjang array s
    copy(c,s) // mengkopi isi dari s = array ke variabel c = array
    fmt.Println("cpy:", c)
    
    l := s[2:5] //mencetak nilai array mulai dari urutan 2 kiri = abc = 0,1,2 dan ke urutan 5 = f sehingga (ab + f) - [a b c d e f] = [c, d, e] berbentuk matriks seperti ini
    fmt.Println("sl1:", l)
    
    l = s[:5] // mencetak urutan array 0 - 4 kecuali urutan 5
    fmt.Println("sl2:", l) 
    
    l = s[2:] // mencetak array mulai dari urutan ke dua yaitu c , urutannya 0 = a b =1 c= 2
    fmt.Println("sl3:", l)
    
     t := []string{"g", "h", "i"} //meng insert nilai array dengan setter value secara langsung
    fmt.Println("dcl:", t)
    
    twoD := make([][]int, 3) // mencetak array multidimensional
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

// $ go run slices.go
// emp: [  ]
// set: [a b c]
// get: c
// len: 3
// apd: [a b c d e f]
// cpy: [a b c d e f]
// sl1: [c d e]
// sl2: [a b c d e]
// sl3: [c d e f]
// dcl: [g h i]
// 2d:  [[0] [1 2] [2 3 4]]
