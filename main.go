package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 {
		// 存在
		return true
	}

	return false
}

func main()  {
	db := InitDB()
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须位11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name,telephone, password)
		//  查询
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
			return
		}
		// 新建
		db.Create(&User{
			Name: name,
			Telephone: telephone,
			Password: password,
		})



		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
	fmt.Println("work")

}
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDB() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	return db

}