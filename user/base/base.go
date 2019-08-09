package base

import (
	"github.com/LiuBaiSMD/goProPlugins/user/base/config"
	"github.com/LiuBaiSMD/goProPlugins/user/user/base/db"
	"github.com/LiuBaiSMD/goProPlugins/user/user/base/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
