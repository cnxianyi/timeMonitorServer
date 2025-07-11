package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"timeMonitorServer/types"
)

var (
	Mdb *gorm.DB // 全局数据库连接池变量，初始值为 nil
)

func InitMysql() error {
	// 读取 MYSQL 环境变量
	dsn := os.Getenv("MYSQL")
	if dsn == "" {
		return fmt.Errorf("MYSQL: 环境变量未设置")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(types.ProcessModel{}, types.TitleModel{}, types.TitleClassModel{})
	if err != nil {
		return err
	}

	Mdb = db

	return nil
}
