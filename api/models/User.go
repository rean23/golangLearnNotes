package models

//表名是结构体名称的复数形式 也就是指向表名为users
type User struct {
	Model
	Name string //列名是字段名的蛇形小写 也就是Name等于name
	Age int
}