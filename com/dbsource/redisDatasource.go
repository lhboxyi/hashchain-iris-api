package datasource
//
//import (
//	red "github.com/gomodule/redigo/redis"
//	"github.com/sirupsen/logrus"
//	"time"
//)
//
//type Redis struct {
//	pool *red.Pool
//}
//
//var redis *Redis
//
////func RedisClient() *Redis {
////	address := viper.GetString("redis.cluster.nodes")
////	fmt.Println("=======",address)
////	redis := new(Redis)
////	redis.pool = &red.Pool{
////		MaxIdle:   viper.GetInt("redis.maxIdle"),
////		MaxActive: 0,
////		//IdleTimeout: time.Duration(120),
////		IdleTimeout: viper.GetDuration("redis.timeout"),
////		Dial: func() (red.Conn, error) {
////			return red.Dial(
////				"tcp",
////				address,
////				red.DialReadTimeout(time.Duration(5000)*time.Millisecond),
////				red.DialWriteTimeout(time.Duration(5000)*time.Millisecond),
////				red.DialConnectTimeout(time.Duration(5000)*time.Millisecond),
////				red.DialDatabase(0),
////				red.DialPassword(viper.GetString("redis.password")),
////			)
////		},
////	}
////	return redis
////}
//
//func InitRedis() {
//	redis = new(Redis)
//	redis.pool = &red.Pool{
//		MaxIdle:     256,
//		MaxActive:   0,
//		IdleTimeout: time.Duration(120),
//		Dial: func() (red.Conn, error) {
//			return red.Dial(
//				"tcp",
//				"10.59.72.73:6379",
//				red.DialReadTimeout(time.Duration(50000)*time.Millisecond),
//				red.DialWriteTimeout(time.Duration(50000)*time.Millisecond),
//				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
//				red.DialDatabase(0),
//				red.DialPassword("Ryh74OCX1u"),
//			)
//		},
//	}
//}
//
//func ExecRedisCommand(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
//	//获取客户端连接
//	con := redis.pool.Get()
//	if err := con.Err(); err != nil {
//		logrus.Errorf("获取redis客户端连接失败:%s", err.Error())
//		return nil, err
//	}
//	defer con.Close()
//	params := make([]interface{}, 0)
//	params = append(params, key)
//
//	if len(args) > 0 {
//		for _, v := range args {
//			params = append(params, v)
//		}
//	}
//	return con.Do(cmd, params...)
//}
//
//
////func ExecRedisCommand(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
////	//获取客户端连接
////	redisClient := RedisClient()
////	con := redisClient.pool.Get()
////	if err := con.Err(); err != nil {
////		logrus.Errorf("获取redis客户端连接失败:%s", err.Error())
////		return nil, err
////	}
////	defer con.Close()
////	params := make([]interface{}, 0)
////	params = append(params, key)
////
////	if len(args) > 0 {
////		for _, v := range args {
////			params = append(params, v)
////		}
////	}
////	return con.Do(cmd, params...)
////}
