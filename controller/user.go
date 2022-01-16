package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goshiyan/common"
	"goshiyan/model"
	"goshiyan/utils"
	"log"
	"net/http"
)

func Register(ctx *gin.Context)  {
	db := common.GetDB()
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
		name = utils.RandomString(10)
	}

	log.Println(name,telephone, password)
	//  查询
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}
	// 新建
	db.Create(&model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	})



	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}



//  控制器扩展函数
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 {
		// 存在
		return true
	}

	return false
}