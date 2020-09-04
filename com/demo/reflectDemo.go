package demo

import (
	"fmt"
	"reflect"
)

func ReflectDemo(obj interface{}) {
	//获取对象的类型和值
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	fmt.Printf("获取对象的类型：%s,值：%v\n", objType, objValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i).Interface()
		fmt.Printf("属性信息【%s: %v = %v】\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		fmt.Printf("方法名【%s】,方法类型【%v】,方法路径【%v】\n", m.Name, m.Type,m.Index)
	}
}
