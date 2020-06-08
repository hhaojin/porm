package porm

import (
	"github.com/hhaojin/porm/Helper"
	"github.com/jmoiron/sqlx"
	"log"
)

var mySqlDB *sqlx.DB

func init() { //先处理mysql
	config, err := Helper.LoadConfig()
	if err != nil {
		log.Println(err)
	}
	db, err := sqlx.Connect(config.DB.Driver, config.DB.DSN)
	if err != nil {
		log.Fatal("DB Connect error:", err)
	}
	mySqlDB = db
}

//获取mysql对应的DB对象
func MySql() *sqlx.DB {
	return mySqlDB
}
