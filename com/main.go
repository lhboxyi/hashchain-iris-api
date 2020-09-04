package main

import (
	"iris-api/com/config"
	datasource "iris-api/com/dbsource"
	"iris-api/com/web/routes"
)

/**
初始化配置函数
*/
func init() {
	config.InitConfig("")
}


func main() {

	/*passwordOK := "admin"
	passwordERR := "adminxx"

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(hash)

	encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePW)

	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

	// 错误密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}*/

	datasource.InstanceMaster()

	//启动iris服务
	routes.Register() //go version




	//dbsource.InitRedis()
	//command, _ := dbsource.ExecRedisCommand("set", "name", "hujianjun675")
	//fmt.Println(command)
}




