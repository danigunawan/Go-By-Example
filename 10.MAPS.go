package main
import "fmt"

func main(){
    
    m := make(map[string]int)
    
    m["k1"] = 7
    m["k2"] = 13
    
    fmt.Println("map:", m) // mencetak key maps yang dibungkus
    
    v1 := m["k1"] // mencetak nilai key maps k1
    fmt.Println("v1:", v1)
    
    fmt.Println("len", len(m)) // cek panjang key maps
    
    delete(m, "k2") //menghapus key maps k2
    fmt.Println("map:", m)
    
    _, prs := m["k2"] //mencek apakah ada array k maps di key yang k2 jika ada true jika tidak false
    fmt.Println("prs:", prs)
    
    n := map[string]int{"foo": 1, "bar": 2} // menginsert langsung nilai key map ke bungkus map dengan value static
    fmt.Println("map:", n)
}

// $ go run maps.go 
// map: map[k1:7 k2:13]
// v1:  7
// len: 2
// map: map[k1:7]
// prs: false
// map: map[foo:1 bar:2]