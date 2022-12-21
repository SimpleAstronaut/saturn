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
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"saturn/api"
)

// Users 读取用户数据的struct
type Users struct {
	Username string
	Password string
}

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
		} else {
			file, err := ioutil.ReadFile("./data/users.json")
			var u Users
			if err != nil {
				fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
				return
			}

			//解码json
			err = json.Unmarshal(file, &u)
			if err != nil {
				log.Fatal("Error during Unmarshal(): ", err)
			}
			c.JSON(200, u)
		}
	})

	r.GET("/getlist", func(c *gin.Context) {
		mode := c.Query("mode")
		list := api.Getlist(mode)
		if list == "mode error" {
			c.HTML(http.StatusOK, "err.html", gin.H{
				"errmsg": "Mode is null",
			})
		} else if list == "error" {
			c.HTML(http.StatusOK, "err.html", gin.H{
				"errmsg": "mode error",
			})
		} else {
			c.JSON(200, list)
		}
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
