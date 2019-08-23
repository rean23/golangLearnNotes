package main

import . "fmt"

/**
定义一个结构体
*/
type person struct {
	name string
	age  int
	sex  string
}

//组合
type student struct {
	person // 此person为匿名字段,在结构体中用person调用
	class  string
	age    int
}

func main() {
	// 定义结构体并赋值
	var p person
	p.name = "rean"
	p.age = 29
	p.sex = "male"

	//另外几种生命方式
	//1.按照顺序提供值
	p1 := person{"sin", 20, "woman"}
	//p1 : {sin 20 woman}
	Println(p1)

	//2.按照kv顺序提供值
	p2 := person{name: "yuna", age: 18, sex: "woman"}
	//p2:{yuna 18 woman}
	Println(p2)

	//3.通过new创建一个指针
	p3 := new(person)
	p3.name = "joker"
	p3.age = 16
	p3.sex = "male"
	//p3 : &{joker 16 male}

	// 赋值组合结构体
	p4 := student{person: person{name: "zacks", age: 18, sex: "male"}, class: "no.1", age: 30}
	//p4 : student{person{"zacks",18,"male"},"no.1"}
	Println(p4.age)//这里优先访问外层的age  如果要访问里层的可以p4.person.age
}
