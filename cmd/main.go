package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Age     string `json:"age"`
	School  string `json:"school"`
	Address string `json:"address"`
	ID      string `json:"id"`
}

func main() {

	// 连接 Mysql
	dsn := "root:passwd@tcp()/myweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	data := User{}
	err = db.AutoMigrate(&data)
	if err != nil {
		return
	}

	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {

		userId := c.Param("id")

		var lists []User

		db.Where("id = ? ", userId).Find(&lists)

		if len(lists) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "学生档案查询失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "学生档案查询成功",
				"data": lists,
			})
		}
	})

	r.POST("/users", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "创建学生档案失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "学生档案建立成功",
				"data": data,
			})
		}
		db.Create(&data)
	})
	err = r.Run(":8090")
	if err != nil {
		log.Fatal("请检查网络连接状态")
	}
}
