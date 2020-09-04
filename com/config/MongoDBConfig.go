package config

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/**
获取mongo客户端连接
*/
func MongoClient() (*mongo.Client, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//调用WithTimeout之后defer cancel()
	//defer cancel()
	mongoUrl := viper.Get("mongo.url").(string)
	mongoPort := viper.Get("mongo.port").(string)
	mongoName := viper.Get("mongo.username").(string)
	mongoPwd := viper.Get("mongo.password").(string)
	database := viper.Get("mongo.database").(string)
	authSource := viper.Get("mongo.authSource").(string)

	uri := "mongodb://" + mongoName + ":" + mongoPwd + "@" + mongoUrl + ":" + mongoPort + "/" + database + "?authSource=" + authSource
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Error("获取mongo连接失败")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Error(err)
	}
	//logrus.Info("获取mongo连接成功")
	return client, ctx
}
