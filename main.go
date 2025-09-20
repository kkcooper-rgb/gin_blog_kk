package main

import (
	"gin_blog_kk/core"
	"gin_blog_kk/global"
	"gin_blog_kk/initialize"
)

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()
	defer global.Redis.Close()
	initialize.InitCron()
	core.RunServer()
}
