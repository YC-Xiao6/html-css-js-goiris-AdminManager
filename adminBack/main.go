package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"goProject/adminBack/config"
	"goProject/adminBack/controller"
)


func main() {
	app := iris.New()
	//解析json配置文件
	if err := config.InitConfig("adminBack\\config.json");err != nil{
		fmt.Printf(err.Error())
	}
	conf := config.ServConfig
	//服务器部署的地址与端口号
	addr := fmt.Sprintf("%s:%d",conf.Ip,conf.Port)
	staticPath := conf.StaticPath
	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析。
	app.RegisterView(iris.HTML(staticPath, ".html"))
	//配置资源文件的访问路径，一参：访问的路径，二参：文件的相对路径
	app.StaticWeb("/","adminBack\\static")
	//创建转发路由配置
	mvc.Configure(app.Party("/"), func(a *mvc.Application) {
		//创建首先访问的路径
		a.Party("/").Handle(new(controller.RouteController))
		a.Party("/admin").Handle(new(controller.AdminController))
	})

	app.Run(iris.Addr(addr))
}
