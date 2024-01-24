package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	var err error

	// 这里你需要用你的实际数据库连接字符串替换下面的dsn
	dsn := "new_executor_qa:Pp8xCfeZ1fjVHDBG1JHeZYi2@tcp(nonprod-serverless-v2-cluster.cluster-cz29fkykottg.ap-southeast-1.rds.amazonaws.com:3306)/new_executor_qa?charset=utf8&parseTime=True&loc=Local" // MySQL connection string
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}

}

func GetDB() *gorm.DB {
	return db
}
