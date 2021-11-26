package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"will/model"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	//user := "root"
	//pass := "123456"
	//dbName := "test"
	//host := "192.168.1.8"
	//port := 3306
	//charset := "utf8mb4"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user,
	//	pass,
	//	host,
	//	port,
	//	dbName,
	//	charset,
	//)

	db, err := gorm.Open(sqlite.Open("will.db"), &gorm.Config{})
	if err != nil {
		panic("mysql connect error:" + err.Error())
	}

	// 迁移 schema
	AutoMigrate(model.Will{}, db)

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate(model interface{}, db *gorm.DB) {
	err := db.AutoMigrate(model)
	if err != nil {
		panic("mysql AutoMigrate error:" + err.Error())
	}

}
