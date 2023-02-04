package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func newDB(config DBConfig, env Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// NOTE: dev環境の時はSQLはすべてログに吐く
	if env == Env_Dev {
		db = db.Debug()
	}
	return db, nil
}
