// Package api get module
// return page or blog
// version 1.0.1
// author SimpleAstronaut
// 2022-12-21

package api

import (
	"fmt"
	"io/ioutil"
)

func Get(mode, name string) string {
	//获取blog接口
	if mode == "blog" {
		blogPath := "./blog/" + name + ".json"
		file, err := ioutil.ReadFile(blogPath)
		if err != nil {
			fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
			return "error"
		}
		result := string(file)
		return result
	} else if mode == "page" {
		pagePath := "./pages/" + name + ".json"
		file, err := ioutil.ReadFile(pagePath)
		if err != nil {
			fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
			return "error"
		}
		result := string(file)
		return result
	} else {
		return "error"
	}
}
