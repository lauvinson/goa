package model

import (
	"singo/util"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

// Casbin
var Enforcer *casbin.Enforcer

func CasbinLoader(connString string) {
	defer func() {
		if recover() != nil {
			util.Log().Panic("连接数据库错误: %s", connString)
			return
		}
	}()
	a := xormadapter.NewAdapter("mysql", connString, true)
	Enforcer = casbin.NewEnforcer("conf/locales/rbac_models.conf", a)
	Enforcer.LoadPolicy()
}
