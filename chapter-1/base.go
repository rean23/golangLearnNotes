/**
Go基础
*/
package main

// 导入多个模块
import (
	"fmt"
)

//定义变量 单个变量 var variableName type  多个变量 var v1,v2,v3 type

var e, f, g = 4, 5, 6 //e,f,g := 4,5,6这样简写  会无法通过编译

func main() {
	// 定义多个变量  自动推导数值类型
	//var a, b, c = 1, 2, 3

	//上面代码的简写
	a, b, c := 1, 2, 3   // 只能在函数内部使用 无法在函数外部使用，所以定义全局变量还是需要var关键字
	fmt.Println(a, b, c) //1 2 3

	fmt.Println(e, f, g) // 4 5 6

	var1, var2, _ := "var1", "var2", "var3" // _表示占位符，表示赋予的值会被丢弃
	fmt.Println(var1, var2)

	// 如果定义变量但未使用则无法通过编译
	//var i int;  错误信息为unused variable i

	// 定义常量 在 Go 程序中，常量可定义为数值、布尔值或字符串等类型。
	const name = "rean"
	const Pi float32 = 3.1415926
	fmt.Println(name, Pi)

	/**
	内置基础类型
	Boolean:在 Go 中，布尔值的类型为 bool，值是 true 或 false，默认为 false。
	*/
	var foo bool     //定义布尔值
	fmt.Println(foo) // false
	/**
	数值类型
	整数类型有无符号和带符号两种。
	Go 同时支持 int 和 uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
	Go 里面也有直接定义好位数的类型：rune, int8, int16, int32, int64 和 byte, uint8, uint16, uint32, uint64。
	其中 rune 是 int32 的别称，byte 是 uint8 的别称。
	*/

	/**
	字符串
	Go 中的字符串都是采用 UTF-8 字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是 string。
	*/
	var firstName = `Jon`
	var lastName = "Snow"
	fmt.Println(firstName, lastName)

	// 字符串是不可变的  比如这样修改:firstName[0] = 'h'

	//正确修改字符串方式
	byteArr := []byte(firstName) //转成byte类型数组
	byteArr[0] = 'H'             //注意 这里是单引号 因为是字符
	//fmt.Printf("%s", string(byteArr)) // Hon

	// 连接字符串
	fmt.Sprintf("%s", firstName+lastName)

	//修改字符串也可以用切片
	s := "H" + firstName[1:]
	fmt.Println(s)

	// 多行字符串
	s = `Hi
Rean`
	fmt.Println(s)

	// 错误类型
	/*err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}*/

	//

	// iota枚举  可以和导入多个包一样的写法 iota只能在常量表达式中使用
	const (
		x = iota
		y = 2
		z // 常量声明省略值时，默认和之前一个值的字面相同。也就是y=iota z=iota
	)

	const v = iota          // 每遇到一个 const 关键字，iota 就会重置，此时 v == 0
	fmt.Println(x, y, z, v) //0 2 2

	const (
		b1 = iota //0
		b2 = "a"  //a
		b3 = iota //2
		b4        //3
		b5        //4
		b6        //5
	)

	//数组
	//var arr [n]type //数组定义方式:n表示数组长度，type表示元素类型

	//定义一个长度为10 的整形数组
	var arr [10]int
	arr[1] = 2
	arr[2] = 3

	/**
	声明了一个长度为3的int数组
	[1 2 3]
	*/
	arr1 := [3]int{1, 2, 3}

	/**
	声明了一个长度为 10 的 int 数组，其中前三个元素初始化为 1、2、3，其它默认为 0
	[1 2 3 0 0 0 0 0 0 0]
	*/
	arr2 := [10]int{1, 2, 3}

	/**
	可以省略长度而采用 `...` 的方式，Go 会自动根据元素个数来计算长度
	[4 5 6]
	*/
	arr3 := [...]int{4, 5, 6}
	fmt.Println(arr1, arr2, arr3)

	// 二位数组
	multiArr := [1][4]int{[4]int{1, 2, 3, 4}}
	//简写
	easyArr := [1][4]int{{1, 2, 3, 4}}
	fmt.Println(multiArr, easyArr)

	// slice 切片
	/**
	注意 slice 和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用 ... 自动计算长度，
	而声明 slice 时，方括号内没有任何字符。
	*/
	sli := []int{1, 2, 3, 4, 5, 6}

	/*
		slice 有一些简便的操作
		slice 的默认开始位置是 0，ar[:n] 等价于 ar[0:n]
		slice 的第二个序列默认是数组的长度，ar[n:] 等价于 ar[n:len(ar)]
		如果从一个数组里面直接获取 slice，可以这样 ar[:]，因为默认第一个序列是 0，第二个是数组的长度，即等价于 ar[0:len(ar)]
		slice还有第三个参数,表示新数组的容量,默认和第二个参数一样
	*/
	fmt.Println(sli[:])

	//mao 映射
	// 生命方式 var numbers map[string]int
	// 声明一个 key 是字符串，值为 int 的字典, 这种方式的声明需要在使用之前使用 make 初始化
	dict := make(map[string]string)
	dict["name"] = "rean"
	dict["age"] = "30"

	//初始化map
	rating := map[string]string{"name": "rean", "age": "30", "sex": "male"}

	// 删除属性
	//delete(rating, "sex")

	// 获取字典值
	sex, ok := rating["sex"]
	if ok {
		fmt.Println("attribute sex is " + sex)
	} else {
		fmt.Println("attribute sex is not exists")
	}

	//map是引用类型

	// make new 知识点
	/**
		make 用于内建类型（map、slice 和 channel）的内存分配。new 用于各种类型的内存分配。
	内建函数 new 本质上说跟其它语言中的同名函数功能一样：new(T) 分配了零值填充的 T 类型的内存空间，并且返回其地址，即一个 *T 类型的值。
	用 Go 的术语说，它返回了一个指针，指向新分配的类型 T 的零值。有一点非常重要：
	new 返回的是指针
	内建函数 make(T, args) 与 new(T) 有着不同的功能，make 只能创建 slice、map 和 channel，并且返回一个有初始值 (非零) 的 T 类型，而不是 *T。
	本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个 slice，是一个包含指向数据（内部 array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice 为 nil。
	对于 slice、map 和 channel 来说，make 初始化了内部的数据结构，填充适当的值。
	make 返回初始化后的（非零）值。
	*/
	m := make(map[string]int)
	m["num"] = 1
	fmt.Println(m)
}
