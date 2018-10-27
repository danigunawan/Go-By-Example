package main
import "fmt"
import "math"

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

//buat action area
func (r rect) area() float64 {
    return r.width * r.height
}

//buat action perim
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius //Pi = 3.14
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4} // inisialisasi declare
    c := circle{radius: 5}
    measure(r) //kita panggil action measure beserta isinya dengan setter argument r
    measure(c)
}

// $ go run interfaces.go
// {3 4}
// 12
// 14
// {5}
// 78.53981633974483
// 31.41592653589793