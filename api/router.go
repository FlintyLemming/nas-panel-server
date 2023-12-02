package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/metrics", getMetricsHandler(db))
	// 其他路由...
}
