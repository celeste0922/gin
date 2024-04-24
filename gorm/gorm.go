package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HelleWord struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

func Gorm() {
	db, err := gorm.Open("mysql", "root:celeste0922@(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
		db.AutoMigrate(&HelleWord{}) //依据结构体创建表

		//新增
		//db.Create(&HelleWord{        //创建明细
		//	Name: "ww",
		//	Sex:  true,
		//	Age:  20,
		//})
		hello := HelleWord{}
		var hello2 []HelleWord

		//查询
		db.Where("name = ?", "yy").Find(&hello2) //可+or
		db.First(&hello, "name = ?", "yy")
		//db.Find(&hello2)

		//修改
		//db.Where("name = ?", "ww").First(&HelleWord{}).Update("name", "qq") //updates

		//删除（软删除）
		//db.Where("name = ?", "yy").Delete(&HelleWord{}) //.Unscoped()物理删除

		//fmt.Println(hello)
		//fmt.Println(hello2)
	}
	defer db.Close()
}
