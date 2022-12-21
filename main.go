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

// Blog Get 读取Blog数据的struct
type Blog struct {
	title  string
	author string
	time   string
	blog   string
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

	//获取列表接口路由
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

	//获取文章接口路由
	r.GET("/blog", func(c *gin.Context) {
		name := c.Query("name")
		blog := api.Get("blog", name)
		if blog == "error" {
			c.HTML(http.StatusOK, "err.html", gin.H{
				"errmsg": "blog error",
			})
		} else {
			var b Blog
			err := json.Unmarshal([]byte(blog), &b)
			if err != nil {
				fmt.Println("json转换失败")
			}

			c.JSON(200, blog)

			//渲染html
			//TODO 无法渲染HTML待修复
			/*c.HTML(http.StatusOK, "page.html", gin.H{
				"pageTitle": b.title,
				"title":     b.title,
				"author":    b.author,
				"time":      b.time,
				"text":      b.blog,
			})*/
		}
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
