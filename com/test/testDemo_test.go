package test

import (
	"fmt"
	"path/filepath"
	"testing"
)

/**
 * 测试文件的文件名需要以_test.go 为结尾，测试用例需要以 TestXxxx 的样式存在。
 */
func TestFileInfo(t *testing.T)  {
	fmt.Println(filepath.Rel("/home/polaris/studygolang", "/home/polaris/studygolang/src/logic/topic.go"))
	fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))
}

//func TestFib(t *testing.T) {
//	var fibTests = []struct {
//		in       int // input
//		expected int // expected result
//	}{
//		{1, 1},
//		{2, 1},
//		{3, 2},
//		{4, 3},
//		{5, 5},
//		{6, 8},
//		{7, 13},
//	}
//
//	for _, tt := range fibTests {
//		actual := Fib(tt.in)
//		if actual != tt.expected {
//			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
//		}
//	}
//}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-1)
}


func Example_GetScore() {
	scores:= []float64{100, 100, 100, 2.1}
	sum:=float64(0)
	for _, score := range scores {
		sum+=score
	}

	fmt.Println(sum)
	// Output:
	// 31.122
}