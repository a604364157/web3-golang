package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * (c.radius * c.radius)
}

/*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
*/
func main() {
	shapes := []Shape{Rectangle{width: 10, height: 5}, Circle{radius: 5}}

	for _, shape := range shapes {
		fmt.Printf("类型: %T\n", shape)
		fmt.Printf("面积: %.2f\n", shape.Area())
		fmt.Printf("周长: %.2f\n\n", shape.Perimeter())
	}
}
