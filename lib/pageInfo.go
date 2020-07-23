package lib

import (
	"math"
)

const (
	//分页每页数据量
	PAGENUM = 10
)

//获取自定义分页信息
func (p *PageStruct) GetPage(total float64, currentPage int, pageInfo *PageStruct, PageNum int) {
	//getPageInfo
	//总条数
	pageInfo.Total = total
	//当前页面
	pageInfo.From = currentPage
	//最后页面
	PageNumToFloat := float64(PageNum)
	pageInfo.LastPage = math.Ceil(float64(pageInfo.Total / PageNumToFloat))
	//每页条数
	pageInfo.PerPage = PageNum
}
