/**
*流程与控制
 */
package main

import "fmt"

//type testInt func(int) bool // 声明了一个函数类型 可以用来约束函数的参数和返回类型

/**
定义模块初始化函数
*/
func init() {
	fmt.Println("package init")
}

/**
模块主函数
*/
func main() {

	// if
	x := 20

	if x > 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//
	/**
	多条件判断
	go的if有一个功能就是可以在其作用域内生命变量
	下面的y只能在该条件逻辑块使用
	*/
	if y := getNum(); y > 20 {
		fmt.Println("The integer is greater than 20")
	} else if y < 10 {
		fmt.Println("The integer is less than 10")
	} else {
		fmt.Println("in else")
	}

	/*for 语法如下:
		for expression1; expression2; expression3 {
	    //...
	}
	*/
	// 和大多数语言一样的for循环
	/*for i:=0;i<=10;i++ {
		fmt.Println(i)
	}*/

	//忽略前后两个赋值表达式
	/*total := 1
	for ;total < 1000; {
		total += total
	}*/
	//上面的简写
	/*total :=1
	for total < 1000 {
		total +=total
	}
	fmt.Println(total)
	*/

	// 遍历映射
	m := map[string]string{"name": "rean", "age": "29", "sex": "male"}
	// 如果不需要k则可以用下划线_站位
	for k, v := range m {
		fmt.Println(k, v)
	}

	// switch流程控制 默认会给每个case后面带有一个break，如果要继续往下执行 则在末尾添加fallthrough关键字
	i := 20
	switch i {
	case 1:
		fmt.Println("i 等于 1")
	case 2, 3, 4:
	case 10:
		fmt.Println("i 在2,3,4,10之中存在!")
	default:
		fmt.Println("i 是一个数字")

	}

	// 获取函数多重返回值
	var1, var2 := foo(3, 4)
	fmt.Println(var1, var2) // 4 3

	bar(4, 3, 2, 1)

	num := 1
	test(&num) // 传递指针到函数

	// defer 延迟执行 defer定义函数是栈结构 先进后出
	/**
	defer 后指定的函数会在函数退出前调用。
	*/
	/*for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 4 3 2 1 0
	}*/

	// 处理错误
	defer func() {
		/**
		recover仅在延迟函数中有效,正常情况下recover()返回nil,在panic抛出致命错误的时候,会获取抛出的错误信息。
		*/
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("error info")
}

func getNum() int {
	return 10
}

/**
定义函数

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    // 这里是处理逻辑代码
    // 返回多个值
    return value1, value2
}
*/

// 指定返回值变量之后,返回的时候不用带上变量名，因为直接在函数里面初始化了
func foo(arg1 int, arg2 int) (type1 int, type2 int) {
	type1 = arg2
	type2 = arg1

	return //若声明了返回变量的话,则return不需要指定变量名
}

// 变参 不定长参数
func bar(args ...int) {
	fmt.Println(args)
}

//指针
// *号表示接受参数的指针
/**
Go 语言中 channel，slice，map 这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。
（注：若函数需改变 slice 的长度，则仍需要取地址传递指针）
*/
func test(a *int) {
	fmt.Println(a)  //内存地址
	fmt.Println(*a) //打印出内存地址上的值
}
