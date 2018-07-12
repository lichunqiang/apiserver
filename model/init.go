package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	cfg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name, true,
		"Local",
	)

	db, err := gorm.Open("mysql", cfg)
	if err != nil {
		panic(err)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//设置最大连接数据, 默认设置为0, 不限制
	//db.DB().SetMaxOpenConns(1000)
	//设置闲置连接数, 开启的连接可复用
	db.DB().SetMaxIdleConns(0)
}

func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),

	)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),

	)
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

//init Database
func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
	//db.Self = GetSelfDB()
	//db.Docker = GetDockerDB()
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
