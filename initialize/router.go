package initialize

import (
	"gin_blog_kk/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(global.Config.System.Env)

	Router := gin.Default()

	return Router
}
