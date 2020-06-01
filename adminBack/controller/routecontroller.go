package controller

import (
	"github.com/kataras/iris/mvc"
)

//创建起始路由路劲
type RouteController struct {}

func (r *RouteController)Get() mvc.Result {
	return mvc.View{
		Name: "login.html",
		Data: map[string]interface{}{
			"Title":"Login Page",
		},
	}
}
func (r *RouteController)GetIndex() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: map[string]interface{}{
			"Title":"Login Page",
		},
	}
}