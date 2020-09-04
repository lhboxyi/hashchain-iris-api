package rpcx

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/client"
)

type ArgsClient struct {
	A int
	B int
}

type ReplyClient struct {
	C int
}

type ArithClient int

var (
	addr = flag.String("addr", "localhost:8787", "server address")
)

func RpcxClient() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &ArgsClient{
		A: 10,
		B: 20,
	}

	for {
		reply := &ReplyClient{}
		err := xclient.Call(context.Background(), "Mul", args, reply) //请求服务函数
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}
func RpcDateClient() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xClient := client.NewXClient("RpcDateServer", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()
	reply := &RpcResult{}
	err := xClient.Call(context.Background(), "GetNowDate", nil, reply) //请求服务函数
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)
}

func RpcBookInfoClient() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xClient := client.NewXClient("RpcDateServer", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()
	reply := &RpcResult{}
	//err := xClient.Call(context.Background(), "GetNowDate", nil, reply) //请求服务函数

	m := &RpcArgs{RpcParam: map[string]interface{}{"id": 1}}

	err := xClient.Call(context.Background(), "GetBookById", m, reply) //请求服务函数
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)
}
