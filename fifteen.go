// package main

// import "fmt"

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	radius float64
// }

// type Parimeter struct {
// 	radius float64
// }

// func (r Circle) Area() float64 {
// 	return 3.14 * r.radius * r.radius
// }

// func (p Parimeter) Area() float64 {
// 	return 2 * 3.14 * p.radius
// }

// func main() {
// 	area := Circle{
// 		radius: 2,
// 	}
// 	pari := Parimeter{
// 		radius: 1,
// 	}
// 	res := area.Area()
// 	fmt.Println("Area of Circle is :: ", res)
// 	res1 := pari.Area()
// 	fmt.Println("Parimeter of Circle is :: ", res1)
// }
