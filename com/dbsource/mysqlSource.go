package datasource

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

/**
 * 获取mysql数据源连接
 * 获取 DB 对象后，连接池是空的，第一个连接在需要的时候才会创建
 */

func MysqlConnectPool() {
	var (
		host     = viper.Get("mysql.host").(string)
		user     = viper.Get("mysql.password").(string)
		password = viper.Get("mysql.username").(string)
		dbName   = viper.Get("mysql.db").(string)
		port     = viper.Get("mysql.port").(int)
	)

	//root:112233@tcp(127.0.0.1:3305)/mygo?charset=utf8
	driveSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", user, password, host, port, dbName)
	mysqlDB, _ := sql.Open("mysql", driveSource)
	mysqlDB.Ping()
}
