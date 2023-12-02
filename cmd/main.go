package main

import (
	"github.com/gin-gonic/gin"
	"nas-panel-server/api"
	"nas-panel-server/config"
	"nas-panel-server/dal"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	db := dal.InitDB(cfg)

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	api.SetupRouter(r, db)

	// 启动服务
	r.Run(cfg.ServerAddress)
}
