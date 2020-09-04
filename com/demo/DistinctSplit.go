package demo

/**
数组切片去重

removeDuplicateElement函数总共初始化两个变量，一个长度为0的slice，一个空map。
由于slice传参是按引用传递，没有创建占用额外的内存空间。
map[string]struct{}{}创建了一个key类型为String值类型为空struct的map，等效于使用make(map[string]struct{})
空struct不占内存空间，使用它来实现我们的函数空间复杂度是最低的。
*/

func DistinctSpiltValue(srcSplit []string) []string {
	//定义个空切片用于存放去重后的值
	distinctSplit := make([]string,0,len(srcSplit))
	//定义一个空map用于存放key值为切片中的值，根据map的指定key获取值是否存在，存在，则不添加
	temp:=make(map[string]struct{})

	for _,item:= range srcSplit{
		if _,ok:=temp[item];!ok{
			temp[item] = struct{}{}
			distinctSplit = append(distinctSplit,item)
		}
	}
	return distinctSplit
}
