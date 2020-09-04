package controller

/**
定义结果返回结构体格式
*/
type Result struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

/**
实例化结果结构体对象
*/
func InstanceResult(code string, objects interface{}, message string) (result *Result) {
	result = &Result{Code: code, Data: objects, Message: message}
	return
}
