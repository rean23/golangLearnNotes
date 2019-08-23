package main

import (
	"fmt"
	"math"
)

/**
面向对象
*/

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 定义method语法:func (r ReceiverType) funcName(parameters) (results)

/**
在使用 method 的时候重要注意几点
虽然 method 的名字一模一样，但是如果接收者不一样，那么 method 就不一样
method 里面可以访问接收者的字段
调用 method 通过 . 访问，就像 struct 里面访问字段一样
类型可以值传递(r Rectangle)也可以引用传递(r *Rectangle)
*/

//定义Rectangle计算面积方法
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Tangle struct {
	Rectangle
}

func (t Tangle) area() float64 {
	return t.width + t.height
}

//自定义类型
type (
	name string
	age  int
	user struct {
		name
		age
	}
)

func main() {

	rectangle := Rectangle{4, 5}
	fmt.Println(rectangle.area())

	circle := Circle{1.22}
	fmt.Println(circle.area())

	fmt.Println(user{"rean", 29})

	// 测试重写方法
	tangle := Tangle{Rectangle{10, 10}}
	fmt.Println(tangle.area())
}
