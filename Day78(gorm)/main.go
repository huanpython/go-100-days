package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model //嵌入常用字段
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/gorm-demo")
	if err != nil {
		panic("failed to connect database")
	}
	//关闭数据库连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&Product{})

	// 插入
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新
	db.Model(&product).Update("Price", 2000)

}
