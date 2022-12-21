// Package api Getlist module
// return blog list or pages list
// version 1.0.0
// author SimpleAstronaut
// 2022-12-21

package api

import (
	"fmt"
	"io/ioutil"
)

func Getlist(mode string) string {
	if mode == "blog" {
		file, err := ioutil.ReadFile("./blog/blog.json")
		if err != nil {
			fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
			return "error"
		}
		result := string(file)
		return result
	} else if mode == "page" {
		file, err := ioutil.ReadFile("./pages/page.json")
		if err != nil {
			fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
			return "error"
		}
		result := string(file)
		return result
	} else {
		return "mode error"
	}
}
