package models

import (
	"github.com/jinzhu/gorm"
	config "api/configs"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	/*ID        uint `gorm:"primary_key"`//主键
	CreatedAt time.Time// 列名为 `created_at`
	UpdatedAt time.Time// 列名为 `updated_at`
	DeletedAt *time.Time `sql:"index"`// 列名为 `deleted_at`
	*/
	gorm.Model
}

var db *gorm.DB

func GetDB() *gorm.DB {

	if db == nil {
		//定义错误
		var err error

		// 连接数据库
		db, err = gorm.Open("mysql", config.GetConfig().Database.Mysql.FormatDSN())

		//有错误则直接抛出
		if err != nil {
			panic(err)
		}

	}

	return db
}
