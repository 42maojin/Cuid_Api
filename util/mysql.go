package util

import (
	"database/sql"
	"fmt"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

// Mysqldb 声明数据库全局变量对象
var Mysqldb *sql.DB

const (
	// DeafultDataBase 默认数据库名字
	DeafultDataBase = "go_test"

	connectCount = 3
)

// InitMysql 初始化数据库
func InitMysql() error {
	datasourceName := mysqlDataSourceName(
		cfg.Mysql.Host,
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Port,
		DeafultDataBase)

	var err error
	for i := 0; i < connectCount; i++ {
		Mysqldb, err = sql.Open("mysql", datasourceName)
		if err != nil {
			// 错误信息
			Logger.Error("mysql connect fail:" + datasourceName)
		}
		Mysqldb.SetMaxOpenConns(100)
		Mysqldb.SetMaxIdleConns(100)
		Mysqldb.SetConnMaxLifetime(2 * 3600)
		err := Mysqldb.Ping()
		if err == nil {
			break
		}
	}
	return err

}

// mysqlDataSourceName mysql连接
func mysqlDataSourceName(host string, user string, password string, port int, dbName string) string {
	fmt.Println(
		user,
		password,
		host,
		port,
		dbName)
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		user,
		password,
		host,
		port,
		dbName)
}
