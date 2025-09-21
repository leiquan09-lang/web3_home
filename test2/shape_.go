package main

import (
	"fmt"
	"math"
)

// 1. 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. 矩形
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. 圆
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     // 匿名组合：Employee“拥有”Person
	EmployeeID int
}

// 方法接收者用组合类型
func (e Employee) PrintInfo() {
	fmt.Printf("EmployeeID:%d Name:%s Age:%d\n",
		e.EmployeeID, e.Name, e.Age)
}

func main() {
	// 接口多态
	var s Shape
	s = Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle Area:", s.Area(), "Perimeter:", s.Perimeter())

	s = Circle{Radius: 5}
	fmt.Println("Circle    Area:", s.Area(), "Perimeter:", s.Perimeter())

	// 组合
	e := Employee{
		Person:     Person{Name: "Bob", Age: 28},
		EmployeeID: 10086,
	}
	e.PrintInfo()
}
