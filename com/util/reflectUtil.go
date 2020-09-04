package util

/**
 * reflect包最核心的两个数据类型我们必须知道，一个是Type，一个是Value
 */
import (
	"container/list"
	"errors"
	"reflect"
	"strings"
	"unsafe"
)

type FieldInfo struct {
	FiledName  string
	FiledType  interface{}
	FiledValue interface{}
}

type MethodInfo struct {
	MethodName string
	MethodType interface{}
}

/**
 * 根据传递的interface对象获取其的属性和方法信息
 */
func GetTypeAndValue(i interface{}) (iType reflect.Type, iValue reflect.Value) {
	//获取interface对象的类型和值的结构体
	iType = reflect.TypeOf(i)
	iValue = reflect.ValueOf(i)
	return
}

/**
 * 根据传递的interface对象获取属性信息
 */
func ReflectFieldList(i interface{}) (res list.List, err error) {
	//获取interface对象的属性列表
	iType, iValue := GetTypeAndValue(i)
	numField := iType.NumField()
	if numField == 0 {
		return res, errors.New("没有对应属性信息")
	}
	fieldList := list.New()
	for i := 0; i < numField; i++ {
		//获取属性结构体
		sField := iType.Field(i)
		//获取value
		vValue := iValue.Field(i).Interface()
		fieldInfo := &FieldInfo{FiledName: sField.Name, FiledType: sField.Type, FiledValue: vValue}
		fieldList.PushBack(fieldInfo)
	}

	//遍历list集合
	//for e := fieldList.Front(); e != nil; e = e.Next() {
	//	v := (e.Value).(*FieldInfo)
	//	fmt.Println(v)
	//}
	return
}

/**
 * 根据传递的interface对象获取方法信息
 */
func ReflectMethodList(i interface{}) (res list.List, err error) {
	//获取interface对象的方法列表
	iType, _ := GetTypeAndValue(i)
	numMethod := iType.NumMethod()
	if numMethod == 0 {
		return res, errors.New("没有对应方法信息")
	}
	methodList := list.New()
	for i := 0; i < numMethod; i++ {
		//获取方法结构体
		sMethod := iType.Method(i)
		methodInfo := &MethodInfo{MethodName: sMethod.Name, MethodType: sMethod.Type}
		methodList.PushBack(methodInfo)
	}
	return
}

/**
 * 动态调用方法
 */
func DynamicCallMethod(obj interface{}, methodName string, params []reflect.Value) []reflect.Value {
	return reflect.ValueOf(obj).MethodByName(methodName).Call(params)
}

/**
 * 解析结构体的tag标签数据，指定属性名和tagKey，获取对应的value
 */
func GetStructTagByFieldAndTagKey(i interface{}, fieldName, tagKey string) (tagValue interface{}, err error) {
	iType, _ := GetTypeAndValue(i)
	numField := iType.NumField()
	if numField == 0 {
		return tagValue, errors.New("没有对应属性信息")
	}
	for i := 0; i < numField; i++ {
		//获取属性结构体
		sField := iType.Field(i)
		if len(fieldName)>0 && strings.EqualFold(sField.Name, fieldName) {
			return sField.Tag.Get(tagKey), nil
		}
	}
	return nil, errors.New("没有获取到指定tag的key对应的值")
}

/**
获取结构体字节大小
*/
func GetByteSize(s struct{}) uintptr {
	return unsafe.Sizeof(s)
}