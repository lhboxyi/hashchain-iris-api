package gormSource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"iris-api/com/models"
	"log"
)

type OwnDb struct {
	*gorm.DB
}

func MysqlClient() *gorm.DB {
	//config.InitConfig("")
	var (
		host     = viper.Get("mysql.host").(string)
		user     = viper.Get("mysql.password").(string)
		password = viper.Get("mysql.username").(string)
		dbName   = viper.Get("mysql.db").(string)
		port     = viper.Get("mysql.port").(int)
	)

	//为了处理time.Time，您需要包括parseTime作为参数
	mysqlUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatal("获取mysql数据库连接失败！")
	}
	log.Println("成功获取数据库连接！！！")

	//db.AutoMigrate(&entity.Stu{})





	// 检查表`users`是否存在
	b := db.HasTable(&models.Stu{})
	fmt.Println(b)
	return db
}

/**
设置自动迁移
*/
func SetAutoMigrate(db *gorm.DB, obj interface{}) {
	// 全局禁用表名复数, 如果设置为true,`User`的默认表名为`user`
	db.SingularTable(true)

	//自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据。生成的表名为：stus
	db.AutoMigrate(obj)

	// 创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Stu{})
}
