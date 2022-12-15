package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"govue.demo/go_web_0/common"
	"govue.demo/go_web_0/dto"
	"govue.demo/go_web_0/model"
	"govue.demo/go_web_0/response"
	"govue.demo/go_web_0/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	passwd := ctx.PostForm("passwd")
	//数据验证
	if len(tel) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "11位！")
		// ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": })
		return
	}
	if len(passwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "最少6位！")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "最少6位！"})
		return
	}
	//如果没有名称给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, tel, passwd)
	//判断参数存在
	if isTeleExist(DB, tel) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不许再次注册！")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "不许再次注册！"})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败！")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密失败！"})
		return
	}
	newUser := model.User{
		Name:     name,
		Tel:      tel,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)
	ctx.JSON(200, gin.H{
		"message": "成功",
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	tel := ctx.PostForm("tel")
	passwd := ctx.PostForm("passwd")
	//数据验证
	if len(tel) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "11位！"})
		return
	}
	if len(passwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "最少6位！"})
		return
	}

	//判断手机号
	var user model.User
	DB.Where("tel = ?", tel).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "用户不存在！"})
		return
	}
	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误！"})
		return
	}
	//发放Token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常！"})
		log.Printf("token generate error: %v!", err)
		return
	}
	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTeleExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
