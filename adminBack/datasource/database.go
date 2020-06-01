package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"goProject/adminBack/config"
)

var (
	Db = connectMysql()
)

//连接数据库
func connectMysql() (*sqlx.DB) {
	//调用配置文件进行解析
	if err := config.InitConfig("D:\\goProject\\adminBack\\config.json");err != nil{
		fmt.Printf(err.Error())
	}
	conf := config.DBConfig
	//qlx连接数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", conf.User, conf.Passwd, conf.Host, conf.Port, conf.Dbname, conf.Charset)
	Db,err := sqlx.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("mysql connect failed,err: [%v]",err.Error())
	}
	return Db
}