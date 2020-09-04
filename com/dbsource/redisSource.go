package datasource

import (
	"fmt"
	//redis2 "github.com/go-redis/redis"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//func GetClusterClient() *redis2.ClusterClient {
//	var client *redis2.ClusterClient
//
//	client = redis2.NewClusterClient(&redis2.ClusterOptions{
//		Addrs: []string{"10.59.72.72:6379", "10.59.72.73:6379"},
//		Password:"Ryh74OCX1u",
//	})
//
//	err := client.Ping().Err()
//	fmt.Println(err)
//	return client
//}

/**
 * 获取redis集群客户端
 */
func GetClusterClient() *redis.ClusterClient {
	//nodes := viper.GetStringSlice("redis.cluster.nodes")
	timeout := viper.GetDuration("redis.timeout")
	readTimeout := viper.GetDuration("redis.readTimeout")
	writeTimeout := viper.GetDuration("redis.writeTimeout")
	maxRetries := viper.GetInt("redis.cluster.max-redirects")
	password := viper.GetString("redis.password")
	fmt.Println("----password---",password)

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		// 填写master主机
		//Addrs: nodes,
		Addrs: []string{"10.59.72.72:6379","10.59.72.73:6379","10.59.72.74:6379"},
		// 设置密码
		Password: password,
		// 设置连接超时
		DialTimeout: timeout,
		// 设置读取超时
		ReadTimeout: readTimeout,
		// 设置写入超时
		WriteTimeout: writeTimeout,
		MaxRetries:   maxRetries,
	})

	// 发送一个ping命令,测试是否通
	if pingErr := clusterClient.Do("ping").Err(); pingErr != nil {
		logrus.Errorf("连接redis集群客户端失败:%s", pingErr.Error())
	}
	return clusterClient
}
