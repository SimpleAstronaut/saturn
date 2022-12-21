// Package api get module
// return page or blog
// version 1.0.1
// author SimpleAstronaut
// 2022-12-21

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Blog struct {
	title  string
	author string
	time   string
	blog   string
}

func Get(mode, name string) {
	//获取blog接口
	if mode == "blog" {
		blogPath := "./blog/" + name + ".json"
		file, err := ioutil.ReadFile(blogPath)
		var b Blog
		if err != nil {
			fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
			return
		}
		err = json.Unmarshal(file, &b)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
	}
}
