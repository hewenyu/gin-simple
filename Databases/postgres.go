package Databases

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ModelWithHistory []interface{}
)

/*
init 初始化
*/
func init() {
	// export POSTGRES_HOST=localhost POSTGRES_USER=gorm POSTGRES_PWD=gorm POSTGRES_DB=gorm POSTGRES_PORT=9920 POSTGRES_SSLMODE=disable
	db_url := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	DB = OpenPG(db_url)

	SetupDatabase(DB)

	return
}

/*
open 链接PG数据库
dsn "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
*/
func OpenPG(dsn string) (db *gorm.DB) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 禁用缓存
	}), &gorm.Config{})

	if err != nil {
		// config.BaseLog.Fatalf("数据库链接失败", err.Error())
	}

	// config.BaseLog.Info("数据库初始化完成")

	return
}

func ResetDatabase(db *gorm.DB) {
	//cleanDatabase(db)
	CleanDatabase(db)
	SetupDatabase(db)
}

/*
cleanDatabase 删除表
*/
func CleanDatabase(db *gorm.DB) {
	db.Migrator().DropTable(ModelWithHistory...)
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(ModelWithHistory...)
}