package initializer

import (
	"github.com/joho/godotenv"
	"goa/cache"
	"goa/configs"
	"goa/model"
	"goa/util"
	"os"
)

// InitDB 初始化配置项
func InitDB() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 读取翻译文件
	if err := configs.LoadLocales(os.Getenv("I18N_MAPPINGS_PATH")); err != nil {
		util.Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 装载Casbin
	model.CasbinLoader(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
