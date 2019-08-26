package configs

import "github.com/go-sql-driver/mysql"

/**
	定义数据库配置结构体
 */
type Connections struct {
	Mysql mysql.Config
}

var Connection = Connections{
	Mysql: mysql.Config{
		User:                 "root",
		Passwd:               "",
		Addr:                 "127.0.0.1",
		DBName:               "test",
		Collation:            "utf8mb4_unicode_ci",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
}
