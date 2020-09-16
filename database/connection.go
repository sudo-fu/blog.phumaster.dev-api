package database

import (
	"fmt"

	"blog.phumaster.dev-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection nil
var (
	Connection *gorm.DB
)

// Initialize nil
func Initialize() (*gorm.DB, error) {
	dbConfig := config.DB()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return conn, err
}
