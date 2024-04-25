package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Student   []Student
}
type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	//多对多
	Teacher []Teacher `gorm:"many2many:student_teacher;"`
}
type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}
type Teacher struct {
	gorm.Model
	TeacherName string
	Student     []Student `gorm:"many2many:student_teacher;"`
}

func ManyToMany() {
	db, _ := gorm.Open("mysql", "root:celeste0922@(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("success")
	//	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{}) //自动建表
	//	i := IDCard{
	//		StudentID: 123456,
	//	}
	//	s := Student{
	//		StudentName: "yy",
	//		IDCard:      i,
	//	}
	//	t := Teacher{
	//		TeacherName: "teacher1",
	//		Student:     []Student{s},
	//	}
	//	c := Class{
	//		ClassName: "class1",
	//		Student:   []Student{s},
	//	}
	//	_ = db.Create(&c).Error
	//	_ = db.Create(&t).Error
	//}
	defer db.Close()
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.Bind(&student)
		db.Create(&student)
	})
	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		_ = db.Preload("IDCard").Preload("Teacher").First(&student, "id=?", id) //预加载
		c.JSON(200, gin.H{
			"student": student,
		})
	})
	r.Run(":8080")
}
