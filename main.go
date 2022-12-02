package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Tel      string `gorm:"varchar(11);not null;unique`
	Password string `gorm:"size:255;not null`
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/reg", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
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
		//如果没有名称给一个随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, tel, passwd)
		//判断参数存在
		if isTeleExist(db, tel) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "不许再次注册！"})
			return
		}
		//创建用户
		newUser := User{
			Name:     name,
			Tel:      tel,
			Password: passwd,
		}
		db.Create(&newUser)
		ctx.JSON(200, gin.H{
			"message": "成功",
		})
	})
	panic(r.Run())
}
func isTeleExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "123.56.19.92"
	port := "3316"
	database := "test"
	username := "test"
	password := "114514"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("error:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
