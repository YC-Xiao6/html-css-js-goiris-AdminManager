package util

import "strconv"

//数据库端分页
func GetPageNum(currentPage,total,pageNum int) (start,end string) {
	if currentPage * pageNum <= total{
		start = strconv.Itoa((currentPage-1)*pageNum)
		end = strconv.Itoa(pageNum)
		return
	}else {
		start = strconv.Itoa((currentPage-1)*pageNum)
		starts :=(currentPage-1)*pageNum
		end = strconv.Itoa(total-starts)
		return
	}
}
