package standardApi

import (
	"fmt"
	"strings"
)

/**
 *  字符串长度；
 * 求子串；
 * 是否存在某个字符或子串；
 * 子串出现的次数（字符串匹配）；
 * 字符串分割（切分）为[]string；
 * 字符串是否有某个前缀或后缀；
 * 字符或子串在字符串中首次出现的位置或最后一次出现的位置；
 * 通过某个字符串将[]string 连接起来；
 * 字符串重复几次；
 * 字符串中子串替换；
 * 大小写转换；
 * Trim 操作；
 */
func StringsDemo() {
	s := "hello,world"
	fmt.Printf("长度【%d】\n",len(s))
	fmt.Printf("子串【%s】\n",s[:3])
	fmt.Printf("是否存在某个字符或子串【%t】\n",strings.Contains(s,"llo"))
	fmt.Printf("字符串分割【%+v】\n",strings.Split(s,","))
	fmt.Printf("字符串是否有某个前缀或后缀【%t】\n",strings.HasPrefix(s,"he"))
	fmt.Printf("字符或子串在字符串中首次出现的位置或最后一次出现的位置【%d】\n",strings.Index(s,","))
	fmt.Printf("通过某个字符串将[]string 连接起来【%s】\n",strings.Join([]string{s,"golang"},"-"))
	fmt.Printf("字符串重复几次【%s】\n",strings.Repeat(s,2))
	fmt.Printf("字符串中子串替换【%s】\n",strings.Replace(s,"wo","Wo",1))
	fmt.Printf("大小写转换【%s】\n",strings.ToUpper(s))
	fmt.Printf("Trim 操作【%s】\n",strings.Trim(s,""))
}
