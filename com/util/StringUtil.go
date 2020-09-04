package util

import (
	"fmt"
	"strings"
)
/**
builder从性能和灵活性上都是最佳

如果能固定字符串长度，则使用sclice byte是最佳选择
 */
func StringBuilderConcat()  {
	sb:= new(strings.Builder)
	sb.WriteString("hu")
	sb.WriteString("jian")
	sb.WriteString("jun")
	fmt.Println(sb)
}

func StringSplitConcat()  {
	sb:= make([]byte,20)
	sb = append(sb,[]byte("hu")...)
	sb = append(sb,[]byte("jian")...)
	sb = append(sb,[]byte("jun")...)
	fmt.Println(string(sb))
}
