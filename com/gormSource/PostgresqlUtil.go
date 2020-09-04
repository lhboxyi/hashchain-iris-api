package gormSource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

func PostgresSqlClient() {
	var (
		host     = viper.Get("postgres.host").(string)
		user     = viper.Get("postgres.password").(string)
		password = viper.Get("postgres.username").(string)
		dbName   = viper.Get("postgres.db").(string)
	)

	pgUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password)
	db, err := gorm.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("获取数据库连接失败！")
	}
	//db.AutoMigrate()
	db.SingularTable(true)
	test := new(Test)

	db.First(&test)
	fmt.Println(test.TName)

}

type Test struct {
	TId int `gorm:t_id`
	TName string `gorm:t_name`
}
