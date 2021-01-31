package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfigurations struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func GetDBConfig() (DatabaseConfigurations, error) {
	var c DatabaseConfigurations
	err := envconfig.Process("DB", &c)

	return c, err
}

func GetGormDB(c *DatabaseConfigurations) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
