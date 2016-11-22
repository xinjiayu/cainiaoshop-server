package db

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

//DB mysql数据库的实例
var DB *sql.DB

func init() {
	log.SetLevel(log.DebugLevel)
	db, err := getDbConnect()
	if err != nil {
		log.Debugf("连接数据库出错:%v", err)
	}
	DB = db
}

func getDbConnect() (*sql.DB, error) {
	return sql.Open("mysql", "root:Hsl58839010@/cainiaoshop")
}
