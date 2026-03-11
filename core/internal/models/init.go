package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

// 初始化数据库操作对象
func Init() *xorm.Engine {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Xorm NewEngine err:%v", err)
		return nil
	}
	return Engine
}
