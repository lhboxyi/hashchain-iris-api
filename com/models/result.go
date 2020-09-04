package models

type Result struct {
	Data   interface{}
	Success string
	Code string
	Msg string
}

var success="true"
var fail = "false"
var resultCode = "0"

//统一返回成功的result
func returnSuccess() *Result  {
	result:=&Result{
		Success: success,
		Msg: "",
		Code: resultCode,
		Data: "",
	}
	return result
}

//统一返回成功的result 带数据
func returnCodeSuccess(msg string,code string,data interface{}) *Result{
	result:=&Result{
		Success: success,
		Msg: msg,
		Code: code,
		Data: data,
	}
	return result
}

func returnFail() *Result  {
	result:=&Result{
		Success: fail,
		Msg: "",
		Code: resultCode,
		Data: "",
	}
	return result
}

func returnCodeFail(msg string,code string,data interface{}) *Result{
	result:=&Result{
		Success: fail,
		Msg: msg,
		Code: code,
		Data: data,
	}
	return result
}






