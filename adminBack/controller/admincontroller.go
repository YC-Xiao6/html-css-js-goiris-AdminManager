package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goProject/adminBack/models"
	"goProject/adminBack/service"
	"strconv"
	"strings"
)

/*定义结构体对象*/
type AdminController struct {
	//创建会话
}
//注册session
var (sess = sessions.New(sessions.Config{Cookie: "mysession"}))

//const ADMIN = "admin"
//转发路由
func (a *AdminController) BeforeActivation( b mvc.BeforeActivation)  {
	//创建监听
	b.Handle("GET","/index","QueryAll")
	b.Handle("GET","/getSession","GetSession")
	b.Handle("GET","/deleteSession","DeleteSession")
	b.Handle("POST","/login","Login")
	b.Handle("GET", "/delete", "DelAdminById")
	b.Handle("GET", "/modifyInfo", "ModifyInfoById")
	b.Handle("GET", "/queryByName", "QueryByName")
	b.Handle("GET", "/queryBySex", "QueryBySex")
	b.Handle("GET", "/queryByDate", "QueryByDate")
	b.Handle("GET", "/queryAllByCreat", "QueryAllByCreat")
	b.Handle("GET", "/queryAllByModify", "QueryAllByModify")
	b.Handle("POST","/addAdmin","AddAdmin")
	b.Handle("POST","/modifyAdmin","ModifyAdmin")
}
/*
*path:/admin/getSession
*获取session值
 */
func (a *AdminController) GetSession (ctx iris.Context) mvc.Result{
	name := sess.Start(ctx).Get("ADMINNAME")
	if name != nil{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":true,
				"name":name,
			},
		}
	}else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":false,
			},
		}
	}
}
/*
*path:/admin/getSession
*退出删除session
 */
func (a *AdminController) DeleteSession (ctx iris.Context) mvc.Result{
	flag := sess.Start(ctx).Delete("ADMINNAME")
	fmt.Println(flag)
	if flag{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":true,
			},
		}
	}else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":false,
			},
		}
	}
}
/*
*path:/admin/add
*添加新数据
 */
func (a *AdminController)AddAdmin (ctx iris.Context) mvc.Result{
	data := &models.Admin{}
	//将传入的json数据放入data中
	if err := ctx.ReadJSON(data);err != nil{
		panic(err.Error())
	}
	flag := service.AddAdmin(data)
	if flag == false {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : false,
				"success": "添加失败",
				"message":"名称重复，请重试！",
			},
		}
	}else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : true,
				"success": "添加成功",
			},
		}
	}
}

/*
* path:/admin/index
* 查询所有
 */
func (a *AdminController)QueryAll(ctx iris.Context)  mvc.Result{
	//	获取路径中的信息
	w := ctx.Request()
	//	获取url的路径信息
	vars := w.URL.Query()
	curr := vars["curr"][0]
	//根据“and”对日期进行分割，返回一个字符串切片
	currentPage, _ := strconv.Atoi(strings.Split(curr,"and")[0])
	pageNum,_ := strconv.Atoi(strings.Split(curr,"and")[1])
	admins, total := service.QueryAll(currentPage,pageNum)
	//设置状态参数
	ctx.StatusCode(iris.StatusOK)
	//将数据以json格式回传给前端
	return mvc.Response{
		//Name: "login.html",
		Object: map[string]interface{}{
			"status" : true,
			"admin": admins,
			"total": total,
		},
	}
}
/*
* path:/admin/queryAllByCreat
* 查询所有
 */
func (a *AdminController)QueryAllByCreat(ctx iris.Context)  mvc.Result{
	//	获取路径中的信息
	w := ctx.Request()
	//	获取url的路径信息
	vars := w.URL.Query()
	curr := vars["curr"][0]
	//根据“and”对日期进行分割，返回一个字符串切片
	currentPage, _ := strconv.Atoi(strings.Split(curr,"and")[0])
	pageNum,_ := strconv.Atoi(strings.Split(curr,"and")[1])
	admins,total := service.QueryAllByCreat(currentPage,pageNum)
	//将数据以json格式回传给前端
	return mvc.Response{
		//Name: "login.html",
		Object: map[string]interface{}{
			"status" : true,
			"admin": admins,
			"total":total,
		},
	}
}
/*
* path:/admin/queryAllByModify
* 查询所有
 */
func (a *AdminController)QueryAllByModify(ctx iris.Context)  mvc.Result{
	//	获取路径中的信息
	w := ctx.Request()
	//	获取url的路径信息
	vars := w.URL.Query()
	curr := vars["curr"][0]
	//根据“and”对日期进行分割，返回一个字符串切片
	currentPage, _ := strconv.Atoi(strings.Split(curr,"and")[0])
	pageNum,_ := strconv.Atoi(strings.Split(curr,"and")[1])
	admins,total := service.QueryAllByModify(currentPage,pageNum)
	//将数据以json格式回传给前端
	return mvc.Response{
		//Name: "login.html",
		Object: map[string]interface{}{
			"status" : true,
			"admin": admins,
			"total":total,
		},
	}
}
/*
* path:/admin/login
* 用户登录
*/
func (a *AdminController)Login(ctx iris.Context) mvc.Result {
	//创建数据类型
	data := &models.LoginAdmin{}
	//将传入的json数据放入data中
	if err := ctx.ReadJSON(data);err != nil{
		panic(err.Error())
	}
	flag := service.Login(data)
	if flag == false {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : false,
				"success": "登录失败",
				"message":"用户名或者密码错误,请重新登录",
			},
		}
	}else {
		sess.Start(ctx).Set("ADMINNAME", data.Name)
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : true,
				"success": "登录成功",
				"message":"登录成功",
			},
		}
	}
		//对admin进行json编码，返还给adminByte
		//adminByte, _ := json.Marshal(admin)
		//将创建一个session{"admin",admin}
		//a.Session.Set(ADMIN, adminByte)

}
/*
*path：/admin/delete
*根据id删除数据
*/
func (a *AdminController)DelAdminById (ctx iris.Context) mvc.Result {
	//获取请求头
	w := ctx.Request()
	//获取请求头中的路径信息
	vars := w.URL.Query()
	//获取id值
	id := vars["id"][0]
	flag := service.DelAdminById(id)
	if(flag){
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : true,
				"message":"删除成功",
			},
		}
	}else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : false,
				"message":"删除失败,请重试！",
			},
		}
	}
}

/*
* path:/admin/modifyAdmin
* 根据索引找要修改的数据
 */
func (a *AdminController) ModifyAdmin (ctx iris.Context) mvc.Result {
	data := &models.Admin{}
	//获取路径中的信息
	w := ctx.Request()
	//获取请求头中的路径信息
	vars := w.URL.Query()
	//获取id值
	strId := vars["id"][0]
	id,_ :=strconv.Atoi(strId)
	//根据id找相应的数据
	//将传入的json数据放入data中
	if err := ctx.ReadJSON(data);err != nil{
		panic(err.Error())
	}
	data.Id = int64(id)
	flag := service.ModifyAdmin(data)
	if flag == false {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : false,
				"success": "修改失败",
				"message":"名称重复，请重试！",
			},
		}
	}else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status" : true,
				"success": "修改成功",
			},
		}
	}
}
/*
* path:/admin/modifyIndex
* 根据索引找要修改的数据
 */
func (a *AdminController) ModifyInfoById (ctx iris.Context) mvc.Result {
	//获取路径中的信息
	w := ctx.Request()
	//获取请求头中的路径信息
	vars := w.URL.Query()
	//获取id值
	id := vars["id"][0]
	//根据id找相应的数据
	admin := service.ModifyInfoById(id)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":true,
			"admin":admin,
		},

	}
}
/*
* path:/admin/queryByName
* 根据名称进行模糊查询
*/
func (a *AdminController) QueryByName (ctx iris.Context) mvc.Result {
//	获取路径中的信息
	w := ctx.Request()
//	获取url的路径信息
	vars := w.URL.Query()
//	获取想要的值
	names := vars["name"][0]
	name := strings.Split(names,"and")[0]

	currentPage, _ := strconv.Atoi(strings.Split(names,"and")[1])

	pageNum,_ := strconv.Atoi(strings.Split(names,"and")[2])
	admins,total := service.QueryByName(name,currentPage,pageNum)
	if admins == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":false,
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":true,
				"admin":admins,
				"total":total,
			},
		}
	}
}
/*
* path:/admin/queryBySex
* 根据名称进行模糊查询
 */
func (a *AdminController) QueryBySex (ctx iris.Context) mvc.Result {
	//	获取路径中的信息
	w := ctx.Request()
	//	获取url的路径信息
	vars := w.URL.Query()
	//	获取想要的值
	sexs := vars["sex"][0]
	sex := strings.Split(sexs,"and")[0]
	currentPage, _ := strconv.Atoi(strings.Split(sexs,"and")[1])
	pageNum,_ := strconv.Atoi(strings.Split(sexs,"and")[2])
	admins,total := service.QueryBySex(sex,currentPage,pageNum)
	if admins == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":false,
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":true,
				"admin":admins,
				"total":total,
			},
		}
	}
}
/*
* path:/admin/queryByBirth
* 根据日期进行模糊查询
 */
func (a *AdminController)QueryByDate(ctx iris.Context) mvc.Result{
//	获取路径中的信息
	w := ctx.Request()
//	获取url的路径信息
	vars := w.URL.Query()
	date := vars["date"][0]
	//根据“and”对日期进行分割，返回一个字符串切片
	upDate := strings.Split(date,"and")[1]
	downDate := strings.Split(date,"and")[0]
	currentPage, _ := strconv.Atoi(strings.Split(date,"and")[2])
	pageNum,_ := strconv.Atoi(strings.Split(date,"and")[3])
	admins,total := service.QueryByBirth(upDate,downDate,currentPage,pageNum)
	if admins == nil {
		return mvc.Response{
		Object: map[string]interface{}{
		"status":false},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":true,
				"admin":admins,
				"total":total,
			},
		}
	}
}