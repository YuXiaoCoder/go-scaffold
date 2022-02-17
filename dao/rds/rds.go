package rds

import (
	"database/sql"
	"errors"
	"fmt"
	"go-scaffold/pkg/configs"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 全局变量
var (
	db    *gorm.DB
	sqlDB *sql.DB
)

// 初始化连接
func Init() (err error) {
	// Data Source Name
	var dsn string
	if configs.AllConfig.RDS.DriverName == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.AllConfig.RDS.MySQL.Username, configs.AllConfig.RDS.MySQL.Password, configs.AllConfig.RDS.MySQL.Host, configs.AllConfig.RDS.MySQL.Port, configs.AllConfig.RDS.MySQL.Database)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if configs.AllConfig.RDS.DriverName == "postgresql" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", configs.AllConfig.RDS.PostgreSQL.Host, configs.AllConfig.RDS.PostgreSQL.Username, configs.AllConfig.RDS.PostgreSQL.Password, configs.AllConfig.RDS.PostgreSQL.Database, configs.AllConfig.RDS.PostgreSQL.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		return errors.New(fmt.Sprintf("unsupported rds driver name: %s", configs.AllConfig.RDS.DriverName))
	}
	// 数据库连接失败
	if err != nil {
		return err
	}

	// 调试模式
	if configs.AllConfig.RDS.Debug {
		db = db.Debug()
	}

	// 连接池
	sqlDB, err = db.DB()
	if err != nil {
		return err
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(configs.AllConfig.RDS.MaxIdleConns)

	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(configs.AllConfig.RDS.MaxOpenConns)

	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(configs.AllConfig.RDS.ConnMaxLifetime) * time.Hour)

	return nil
}

// 关闭连接
func Close() {
	// https://gorm.io/docs/generic_interface.html
	_ = sqlDB.Close()
}
