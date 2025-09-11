package global

import (
	"gin_blog_kk/config"

	"go.uber.org/zap"
)

var (
	Config *config.Config
	Log    *zap.Logger
)
