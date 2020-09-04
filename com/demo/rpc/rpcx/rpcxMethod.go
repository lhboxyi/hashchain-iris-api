package rpcx

import (
	"context"
	"iris-api/com/service"
	"iris-api/com/util"
)

type RpcArgs struct {
	RpcParam map[string]interface{}
}

type RpcResult struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

type RpcRegisterTypeName string

/**
 * 格式化当前日期rpc服务
 */
func (rpc *RpcRegisterTypeName) GetNowDate(ctx context.Context, rpcArgs *RpcArgs, rpcResult *RpcResult) error {
	nowDate := util.GetTodayTime()
	rpcResult.Code = 200
	rpcResult.Msg = "获取当前格式化时间成功"
	rpcResult.Data = map[string]interface{}{"nowDate": nowDate}
	return nil
}

/**
 * 根据book_id查询book的信息rpc服务
 */
func (rpc *RpcRegisterTypeName) GetBookById(ctx context.Context, rpcArgs *RpcArgs, rpcResult *RpcResult) error {
	service := service.NewBookService()
	id := int(rpcArgs.RpcParam["id"].(int64))
	res := service.Get(id)

	rpcResult.Code = 200
	rpcResult.Msg = "根据书本id获取信息"
	rpcResult.Data = map[string]interface{}{"bookInfo": *res}
	return nil
}
