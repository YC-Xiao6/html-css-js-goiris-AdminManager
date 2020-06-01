package service

import (
	"crypto/md5"
	"fmt"
	"goProject/adminBack/datasource"
	"goProject/adminBack/models"
	"goProject/adminBack/util"
	"strconv"
	"time"
)

var Db = datasource.Db
//添加分页功能
//获取所有用户信息
//引用传递,名称需大写，才能被外包访问
func QueryAll(currentPage, pageNum int) (admins []models.Admin,total int) {
	errs := Db.Get(&total,"SELECT count(*) FROM admin")
	start,end := util.GetPageNum(currentPage,total,pageNum)
	//用较为低级的Query语句进行查询
	str := "select * from admin limit "+start+","+end
	rows,err := Db.Query(str)
	if err != nil || errs!=nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		fmt.Println(errs.Error())
		return
	}
	for rows.Next() {
		//定义对象接收查询数据
		var admin models.Admin
		//写入数据
		err := rows.Scan(&admin.Id,&admin.Name,&admin.Passwd,&admin.Birthday,&admin.Sex,&admin.Phone,&admin.Email,&admin.Addr,&admin.Education,&admin.CreateTime,&admin.LastTime)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		admins = append(admins,admin)
	}
	//关闭结果集（释放连接）
	rows.Close()
	return
}
//根据创建时间先后排序
func QueryAllByCreat(currentPage,pageNum int)(admins []models.Admin,total int){
	errs := Db.Get(&total,"SELECT count(*) FROM admin")
	start,end := util.GetPageNum(currentPage,total,pageNum)
	err := Db.Select(&admins,"select * from admin order by create_time DESC limit ?,?",start,end)
	if err != nil||errs != nil{
		fmt.Println(err.Error())
		return
	}else {
		return
	}
}
//根据创建时间先后排序
func QueryAllByModify(currentPage,pageNum int)(admins []models.Admin,total int){
	errs := Db.Get(&total,"SELECT count(*) FROM admin")
	start,end := util.GetPageNum(currentPage,total,pageNum)
	err := Db.Select(&admins,"select * from admin order by last_time DESC limit ?,?",start,end)

	if err != nil || errs != nil{
		fmt.Println(err.Error(),errs.Error())
		return
	}else {
		return
	}
}
//按姓名查询
func QueryBySex(sex string,currentPage,pageNum int)(admins []models.Admin,total int) {
	//用sqlx获取切片
	errs := Db.Get(&total,"SELECT count(*) FROM admin where sex = ?",sex)
	start,end := util.GetPageNum(currentPage,total,pageNum)
	err := Db.Select(&admins,"select * from admin where sex = ? limit ?,?",sex,start,end)
	if err !=nil || errs != nil{
		fmt.Println(err.Error(),errs.Error())
		return
	}else {
		return
	}
}
//模糊查询用户名
func QueryByName(name string,currentPage,pageNum int) (admins []models.Admin,total int) {
	//select查询后获取切片
	errs := Db.Get(&total,"SELECT count(*) FROM admin where name like '%"+name+"%'")
	start,end := util.GetPageNum(currentPage,total,pageNum)
	err := Db.Select(&admins,"SELECT * FROM admin where name like '%"+name+"%' limit ?,?",start,end)
	if err != nil || errs != nil {
		fmt.Printf("query faied, error:[%v],[%v]", err.Error(),errs.Error())
		return
	}else {
		return
	}
}
//根据date查找相应的
func QueryByBirth(upDate string,downDate string,currentPage,pageNum int) (admins []models.Admin,total int)  {
	if upDate == "" {
		//将字符串形式的日期转化为date
		downDates,_ := time.Parse("2006-01-02",downDate)
		errs := Db.Get(&total,"SELECT count(*) FROM admin where birthday >= ?",downDates)
		start,end := util.GetPageNum(currentPage,total,pageNum)
		err := Db.Select(&admins,"select * from admin where birthday >= ? limit ?,?",downDates,start,end)
		if err != nil||errs != nil{
			fmt.Println(err.Error(),errs.Error())
			return
		}else {
			return
		}
	}
	if downDate == ""{
		upDates,_ := time.Parse("2006-01-02",upDate)
		errs := Db.Get(&total,"SELECT count(*) FROM admin where birthday <= ? ",upDates)
		start,end := util.GetPageNum(currentPage,total,pageNum)
		fmt.Println(start,end)
		err := Db.Select(&admins,"select * from admin where birthday <= ? limit ?,?",upDates,start,end)
		if err != nil||errs != nil{
			fmt.Println(err.Error(),errs.Error())
			return
		}else {
			return
		}
	}else {
		upDates,_ := time.Parse("2006-01-02",upDate)
		downDates,_ := time.Parse("2006-01-02",downDate)
		errs := Db.Get(&total,"SELECT count(*) FROM admin where birthday >= ? and birthday <= ?",downDates,upDates)
		start,end := util.GetPageNum(currentPage,total,pageNum)
		fmt.Println(start,end)
		err := Db.Select(&admins,"select * from admin where birthday >= ? and birthday <= ? limit ?,?",downDates,upDates,start,end)
		if err != nil||errs != nil{
			fmt.Println(err.Error(),errs.Error())
			return
		}else {
			return
		}
	}
}


//用户登录服务
func Login(loginadmin *models.LoginAdmin) bool {
	var admin models.Admin
	//对传进来的密码进行md5加密
	//w.Sum(nil)将w的hash转成[]byte格式
	data:=[]byte(loginadmin.Passwd)
	has:=md5.Sum(data)
	md5Passwd :=fmt.Sprintf("%x",has)
	//将str写入到w中
	//queryStr := fmt.Sprintf("select * from admin where name = %s AND passwd = %s",loginadmin.Name,md5Passwd)
	err := Db.Get(&admin,"select * from admin where name = ? AND passwd =?",loginadmin.Name,md5Passwd)
	if err != nil {
		fmt.Println("get data failed, error:[%v]", err.Error())
		return false
	}
	return true
}
//删除指定用户
func DelAdminById(id string) bool{
	result,err := Db.Exec("delete from admin where id = ?",id)
	if err != nil{
		panic(err.Error())
		return false
	}
	num, _ := result.RowsAffected()
	if num == 0{
		return false
	}
	return true
}
//根据id找相应的信息
func ModifyInfoById(id string) models.Admin  {
	var admin models.Admin
	err := Db.Get(&admin,"select * from admin where id = ?",id)
	if err != nil {
		fmt.Println("get data failed, error:[%v]", err.Error())
	}
	return admin
}
//根据name查找响应用户
func GetAdminByName(name string) bool {
	var admin models.Admin
	err := Db.Get(&admin,"select * from admin where name = ? ",name)
	if err != nil {
		return true
	}
	return false
}
//根据id查找对应的名字
func GetNameById(id string)  string {
	var name string
	err := Db.Get(&name,"select name from admin where id = ? ",id)
	if err != nil {
		fmt.Println(err.Error())
		return name
	}
	return name
}
//添加用户
func AddAdmin(admin *models.Admin) bool {
	if (GetAdminByName(admin.Name)){
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		//对传进来的密码进行md5加密
		//w.Sum(nil)将w的hash转成[]byte格式
		data:=[]byte(admin.Passwd)
		has:=md5.Sum(data)
		md5Passwd :=fmt.Sprintf("%x",has)
		str := "insert into admin (name,passwd,birthday,sex,phone,email,addr,education,create_time,last_time) values (?,?,?,?,?,?,?,?,?,?)"
		result := Db.MustExec(str,admin.Name,md5Passwd,admin.Birthday,admin.Sex,admin.Phone,admin.Email,admin.Addr,admin.Education,timeNow,timeNow)
		_,err := result.LastInsertId()
		if err != nil{
			fmt.Println(err.Error())
			return false
		}else {
			return true
		}
	}else {return false}
}
//修改用户
func ModifyAdmin(admin *models.Admin) bool {
	//判断名字是否是之前的名字
	flag := false
	if GetNameById(strconv.Itoa(int(admin.Id)))==admin.Name{
		flag = true
	}
	if GetAdminByName(admin.Name)||flag{
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		//对传进来的密码进行md5加密
		//w.Sum(nil)将w的hash转成[]byte格式
		data:=[]byte(admin.Passwd)
		has:=md5.Sum(data)
		md5Passwd :=fmt.Sprintf("%x",has)
		str := "update admin set name=?,passwd=?,birthday=?,sex=?,phone=?,email=?,addr=?,education=?,last_time=? where id = ? "
		//输出一个错误信息，一个结果
		_,err := Db.Exec(str,admin.Name,md5Passwd,admin.Birthday,admin.Sex,admin.Phone,admin.Email,admin.Addr,admin.Education,timeNow,admin.Id)
		if err != nil{
			fmt.Println(err.Error())
			return false
		}else {
			return true
		}
	}else {return false}
}
