package thriftServer

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/sirupsen/logrus"
	"iris-api/com/demo/rpc/thriftDemo/rpc"
	"iris-api/com/util"
	"os"
)

const (
	NetworkAddr = "10.59.79.12:19090"
)

type RpcServiceImpl struct {
}

func (t *RpcServiceImpl) FunCall(/*ctx context.Context, */callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	for k, v := range paramMap {
		r = append(r, k+v)
	}
	return
}

func (t *RpcServiceImpl) FunCall2(ctx context.Context, callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	nowDate := util.GetTodayTime()
	paramMap["nowDate"] = nowDate
	return
}

func ThriftRpcServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		logrus.Error("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	logrus.Info("thrift server in", NetworkAddr)
	server.Serve()
}