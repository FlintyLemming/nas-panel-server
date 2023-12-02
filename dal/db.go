package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nas-panel-server/config"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 数据库迁移...
	return db
}
