package main

/**
导入import
*/

/**
两种加载路径写法
相对路径
import "./model" // 当前文件同一目录的 model 目录，但是不建议这种方式来 import
绝对路径
import "shorturl/model" // 加载 gopath/src/shorturl/model 模块
*/
import . "fmt"       //点语法可以省略包前缀 也就是fmt.Println可以省略成Println
import o "os"        //起别名
import _ "./Base.go" //_操作是引入该包，而不直接使用包里面的函数，而是调用了该包里面的 init 函数

func main() {
	Println(o.Environ())
}
