namespace go goDemo.rpc
namespace java javaDemo.rpc

// 测试服务
service ThriftService {
	// 发起远程调用
	list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
	//获取本地ip
	string getLocalIp(),

	//获取当前格式化时间
    string getNowDate(1:string format),
}