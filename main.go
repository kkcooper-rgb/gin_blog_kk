package main

import (
	"gin_blog_kk/core"
	"gin_blog_kk/global"
)

func main() {
	global.Config = core.InitConfig()
}
