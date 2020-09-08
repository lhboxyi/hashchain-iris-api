package util_common

var success="true"
var fail = "false"
var resultCode = "0"

//统一返回成功的result
func returnSuccess() *Dto  {
	result:=&Dto{
		Success: success,
		Msg: "",
		Code: resultCode,
		Data: "",
	}
	return result
}

//统一返回成功的result 带数据
func returnCodeSuccess(msg string,code string,data interface{}) *Dto{
	result:=&Dto{
		Success: success,
		Msg: msg,
		Code: code,
		Data: data,
	}
	return result
}

func returnFail() *Dto  {
	result:=&Dto{
		Success: fail,
		Msg: "",
		Code: resultCode,
		Data: "",
	}
	return result
}

func returnCodeFail(msg string,code string,data interface{}) *Dto{
	result:=&Dto{
		Success: fail,
		Msg: msg,
		Code: code,
		Data: data,
	}
	return result
}
