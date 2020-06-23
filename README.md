# html-css-js-goiris-AdminManager
  基于html-css-js-goiris 的用户管理系统
## 用户信息管理系统
    第一次用go搭建的小项目，用于对一些信息数据的管理操作。
    目前实现的功能：
        1.用户的登录退出，session的创建删除
        2.用户信息管理页面
        3.用户的增加、修改、删除
        4.用户信息按名称模糊查询、按日期期间查询、按性别精准查询
### 项目环境与技术
    JDK：goland
    SDK：GO 1.14.3
    MODULES: go.mod
    前端：HTML\CSS\JS，Layui（只用在分页）
    后端：GO\Iris
    数据库：sqlx
### 文件结构
    adminBack
        ---config
            ---config.go
        ---controller
            ---admincontroller.go
            ---routecontroller.go
        ---datasource
            ---database.go
        ---models
            ---admin.go
            ---result.go
        ---service
            ---adminservice.go
        ---util
            page.go
        ---static
            ---css
            ---js
            ---img
            ...
        config.json
        main.go
