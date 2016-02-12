package main

import (
	"fmt"
    "math"
)

type ssdrecord struct {
    brand string
    prod_name string
    count int

}

type geometry interface{
    area() float64
    perimeter() float64
}

type rect struct {
    width,height float64
}

type circle struct{
    radius float64
}


func (r1 rect) area() float64 {
    return r1.width * r1.height
}

func (r1 rect) perimeter() float64 {
    return 2*(r1.width + r1.height)
}

func (c1 circle) area() float64 {
    return math.Pi * c1.radius * c1.radius
 
}

func (c1 circle) perimeter() float64 {
    return 2 * math.Pi * c1.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println("Measure function > Area: ", g.area())
    fmt.Println("Measure function > Perimeter: ", g.perimeter())
}


func plus(a, b int) int {
	return a + b
}

func sums(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
	result := plus(10, 3)
	fmt.Println("Plus function returns ", result)

	set2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sums(set2...)

    r := ssdrecord{"samsung","evo850",480}
    fmt.Println(r.brand, r.prod_name, r.count)

    r2 := &r
    fmt.Println(r2.brand, r2.prod_name, r2.count)
    fmt.Println(r.brand, r.prod_name, r.count)

    rt := rect{width: 50, height: 20}
    cc := circle{radius: 10}

    measure(rt)
    measure(cc)


}

