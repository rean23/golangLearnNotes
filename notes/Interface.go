package main

import (
	"fmt"
	"reflect"
)

/**
	接口
 */
type Base interface {
	Study() string
	SayHi() string
}

// interface也可以嵌入
type Common interface {
	Base
}

type (
	Male struct {
		name string
		age int
	}
	Woman struct {
		name string
		age int
	}
)

/**
	实现了Base接口里的方法就等于继承了接口
 */
func (m Male) SayHi() string {
	return "Hi,I'm "+m.name
}
func (m Male) Study() string {
	return m.name+" is study"
}

func main() {
	//interface 是一组 method 签名的组合，我们通过 interface 来定义对象的一组行为。

	//interface{}类型的变量可以接受任何值
	/*var a interface{}
	a = 1*/
	var b Base

	t := reflect.TypeOf(b)
	fmt.Println(t.Elem())
}
