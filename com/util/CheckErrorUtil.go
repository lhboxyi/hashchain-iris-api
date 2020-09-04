package util

import (
	"fmt"
	"log"
	"runtime"
)

/**
处理错误检查工具类（可以查找错误文件地址和错误行号）
*/

func CheckError(err error) bool {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			errMessage := fmt.Sprintf("file:%s, line:%d, error:%s", file, line, err.Error())
			log.Print(errMessage)
		} else {
			log.Print(err)
		}
		return true
	}
	return false
}
