package db

import (
	"auth-gateway/cache"
	"auth-gateway/configs"
	"auth-gateway/util/log"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
)

// InitDB 初始化配置项
func InitDB() {
	// 读取翻译文件
	if err := configs.LoadLocales(os.Getenv("I18N_MAPPINGS_PATH")); err != nil {
		log.Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	Database()
	// 装载Casbin
	CasbinLoader()
	cache.Redis()
}

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database() {
	db, err := gorm.Open(os.Getenv("XORM_DRIVER_NAME"), os.Getenv("DATABASE_DSN"))
	// Error
	if err != nil {
		log.Panic("连接数据库不成功", err)
	}
	db.LogMode(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置库配置
	db.InstantSet("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	DB = db
}

// Enforcer Casbin装载器
var Enforcer *casbin.Enforcer

// CasbinLoader casbin配置加载
func CasbinLoader() {
	defer func() {
		if recover() != nil {
			log.Panic("连接数据库错误: %s", os.Getenv("DATABASE_DSN"))
			return
		}
	}()
	a := xormadapter.NewAdapter(os.Getenv("XORM_DRIVER_NAME"), os.Getenv("DATABASE_DSN"), true)
	Enforcer = casbin.NewEnforcer(os.Getenv("CASBIN_RBAC_MODELS_CONF_PATH"), a)
	_ = Enforcer.LoadPolicy()
}
