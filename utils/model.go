package utils

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Viper.GetString("db.User"),
		Viper.GetString("db.Password"),
		Viper.GetString("db.Host"),
		Viper.GetString("db.DbName"))
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql connect err: %v", err)
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Second * 10)
	DB.SetMaxIdleConns(Viper.GetInt("db.MaxIdleConn")) //用于设置闲置的连接数。
	DB.SetMaxOpenConns(Viper.GetInt("db.MaxActiveConn")) //用于设置最大打开的连接数，默认值为0表示不限制。
}

