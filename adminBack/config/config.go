package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
//主程序服务的配置
type AppConfig struct {
	AppName string `json:"app_name"`
	Port int `json:"port"`
	Ip string `json:"ip"`
	StaticPath string `json:"static_path"`
	Mode string `json:"mode"`
}
//数据库连接的配置
type Db struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Dbname string `json:"dbname"`
	User string `json:"user"`
	Passwd string `json:"passwd"`
	Charset string `json:"charset"`
}
//封装整个配置文件
type BaseConfig struct {
	AppConfig
	Db `json:"dbConfig"`
}
//定义两个配置的地址
var (
	ServConfig *AppConfig
	DBConfig *Db
)

//初始化服务器配置
func InitConfig(filename string)(err error) {
	var (
		content []byte
		conf BaseConfig
		)
	//是io的进化版，方便io的读取
	content,err =ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
		return
	}
	//解析json编码的数据并将结果存入conf指向的值。
	err = json.Unmarshal(content,&conf)
	if err != nil{
		fmt.Println(err)
		return
	}
	ServConfig = &conf.AppConfig
	DBConfig = &conf.Db
	return
}
