/*
 * The Saturn Project
 * S-Blog v2.0
 * based on golang Gin
 * Author: SimpleAstronaut
 * Version: v1.0.0_alpha_22122001
 * 2022-12-20
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")

	//主页面测试路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//登录获取token接口
	r.GET("/login", func(c *gin.Context) {
		username := c.DefaultQuery("username", "null")
		password := c.DefaultQuery("password", "null")

		//判断username和password是否为空
		if username == "null" || password == "null" {
			c.HTML(http.StatusOK, "err.html", gin.H{
				"errmsg": "Username or Password is null",
			})
		}
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
