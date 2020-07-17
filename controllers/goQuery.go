package controllers

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"offergo/lib"
	"strings"
)

type GoQueryController struct {
	baseController
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

func (g *GoQueryController) SearchDocument() {
	url := g.Input().Get("url")
	//必传参数验证
	validation := map[string]interface{}{
		"url": url,
	}
	g.requestFilter(validation)

	// 解析html
	res, err := http.Get(url)
	if err != nil {
		g.responseError(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		g.responseError(res.Status)
	}
	// 读取dom树
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		g.responseError(err.Error())
	}
	var html string
	var img []string
	// 抓取对应的html
	doc.Find(".rich_media_content").Each(func(i int, s *goquery.Selection) {
		html, _ = s.Html()
	})
	//去除换行
	html = strings.Replace(html, "\n", "", -1)

	//抓取对应的img地址
	doc.Find(".rich_media_content section section section section img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("data-src")
		img = append(img, src)
	})
	var result lib.SearchDocumentStruct
	result.Html = html
	result.Img = img
	g.responseSuccess(result)
}

func (g *GoQueryController) ReplaceSearchDocument() {
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
	doc.Find(".rich_media_content section section section section img").Each(func(i int, s *goquery.Selection) {
		_, err := s.Attr("data-src")
		if err {
			s.SetAttr("src", replaceSlice[key])
			key++
		}
	})

	result, err := doc.Find(".rich_media_content").Html()
	if err != nil {
		g.responseError(err.Error())
	}
	g.responseSuccess(result)
}
