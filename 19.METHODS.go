package main
import "fmt"

type rect struct{
    width, height int
}

// kita buat action untuk area dengan deklarenya direturn dari base structs 
func (r * rect) area() int{ //* rect itu receiver base dari structs yang nantinya dipake pada var r di method area() dipanggil dengan dot (.)
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}


// $ go run methods.go 
// area:  50
// perim: 30
// area:  50
// perim: 30
