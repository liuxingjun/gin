package lib

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Gorm *gorm.DB

func init() {
	dbConfig := mysql.NewConfig()
	dbConfig.User = "root"
	dbConfig.Passwd = "123"
	dbConfig.Net = "tcp"
	dbConfig.Addr = "192.168.0.222:3306"
	dbConfig.DBName = "gin"

	//https://github.com/go-sql-driver/mysql#parameters
	dbConfig.Params = map[string]string{
		//"charset":"utf8mb4", //不建议使用charset参数，因为它会向服务器发出额外的查询
		"parseTime": "true",
		"loc":       "Asia/Shanghai",
	}
	var err error
	Gorm, err = gorm.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	Gorm.LogMode(true)
}
