package demo

import "fmt"

/**
下面拷贝intArr的值到指针切片中，获取到内容都是最后一个值
原因：for range内部还是for循环，再遍历内容时，每次遍历的v都是对同一个元素的遍历赋值。也就是说如果直接对v取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值
解决方法只需对遍历的值进行局部定义即可
*/
func ForRangeDemo1() {
	intArr := [3]int{1, 2, 3}
	var res1 []*int
	var res2 []*int
	var res3 []*int
	for k, v := range intArr {
		res1 = append(res1, &v)

		//更改方法一
		vTmp := v
		res2 = append(res2, &vTmp)

		//更改方法二
		res3 = append(res3, &intArr[k])
	}
	fmt.Println("拷贝切片结果1:", *res1[0], *res1[1], *res1[2]) //拷贝切片结果 3 3 3
	fmt.Println("拷贝切片结果2:", *res2[0], *res2[1], *res2[2]) //拷贝切片结果 1 2 3
	fmt.Println("拷贝切片结果3:", *res3[0], *res3[1], *res3[2]) //拷贝切片结果 1 2 3
}

/**
会停止，因为遍历v切片前，先做了拷贝，所以遍历期间对原来v的修改不会反映到遍历中
 */
func ForRangeDemo2() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println(v)
}

/**
大数组这样遍历有啥问题？
遍历前的拷贝对内存是极大浪费

优化：

对数组取地址遍历
for i, n := range &arr
对数组做切片引用
for i, n := range arr[:]

 */
func ForRangeDemo3() {
	//假设值都为1，这里只赋值3个
	var arr = [102400]int{1, 1, 1}
	for i, n := range arr {
		//just ignore i and n for simplify the example
		_ = i
		_ = n
	}
}