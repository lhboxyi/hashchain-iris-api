package rpcx

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/spf13/viper"
)

/**
 * 注册服务函数服务端
 */
func RpcxServer() {
	s := server.NewServer()
	//注册服务函数
	//s.RegisterName("Arith", new(Arith), "")
	s.RegisterName("RpcDateServer", new(RpcRegisterTypeName), "")

	//启动rpcx服务
	fmt.Println("rpc服务地址为:", fmt.Sprintf("%s:%d", viper.GetString("rpcx.host"), viper.GetInt32("rpcx.port")))
	s.Serve("tcp", fmt.Sprintf("%s:%d", viper.GetString("rpcx.host"), viper.GetInt32("rpcx.port")))

}

