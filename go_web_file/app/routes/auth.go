package routes

import (
	//"bytes"
	"filrserver/app/middlewares"
	"filrserver/pkgs/config"

	//"io/ioutil"
	"net/http"

	"filrserver/pkgs/model"

	"github.com/gin-gonic/gin"
)

var user model.User

func authHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	user = model.User{}

	//同时支持json和from-data
	//zlog.SugLog.Info(c.PostForm("username"))～
	//err := c.ShouldBindBodyWith(&user, binding.JSON)
	//if err != nil {
	//	//user.Username = c.PostForm("username")
	//	a, _ := c.Get(gin.BodyBytesKey)
	//	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(a.([]byte)))
	//	//user.Password = c.PostForm("password")
	//	zlog.SugLog.Info(c.GetPostForm("password··"))
	//
	//}
	c.ShouldBindJSON(&user)
	if user == (model.User{}) {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return

	}
	// 校验用户名和密码是否正确
	if chek_user(user.Username, user.Password) {
		// 生成Token
		tokenString, _ := middlewares.CreateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"token":    tokenString,
				"username": user.Username,
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "用户名或密码错误",
	})
	return
}

func chek_user(username string, password string) bool {
	for _, u := range config.Users.Users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
