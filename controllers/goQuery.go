package controllers

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type GoQueryController struct {
	baseController
	result
}

//反序列化存放结构体
var replaceSlice []string

func (g *GoQueryController) ReplaceDocument() {
	url := g.Input().Get("url")
	replaceArray := g.Input().Get("replaceArray")
	//必传参数验证
	validation := map[string]interface{}{
		"url":          url,
		"replaceArray": replaceArray,
	}
	g.requestFilter(validation)

	// 解析html
	res, err := http.Get(url)
	if err != nil {
		g.responseError(err.Error())
	}
	defer res.Body.Close()
	//解析replaceArray存放进切片里面
	json.Unmarshal([]byte(replaceArray), &replaceSlice)
	if res.StatusCode != 200 {
		g.responseError(res.Status)
	}
	// 读取dom树
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		g.responseError(err.Error())
	}
	key := 0
	// 查找img值并替换src
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		_, err := s.Attr("data-src")
		if err {
			s.SetAttr("src", replaceSlice[key])
			key++
		}
	})

	result, err := doc.Html()
	if err != nil {
		g.responseError(err.Error())
	}
	//剔除goquery添加的html/head/body标签
	result = strings.Replace(strings.Replace(result, "<html><head></head><body>", "", -1), "</body>", "", -1)
	g.responseSuccess(result)
}
