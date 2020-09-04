package standardApi

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	//可以通过读取标准输入
	bytes := make([]byte, num)

	if readNum, err := reader.Read(bytes); err == nil {
		return bytes[:readNum], err
	}
	return bytes, nil
}

/**
 * 获取Read对象的几种方式
 * 标准输入 os.Stdin
 * 文件对象
 * 读取字符串 strings.NewReader("hujianjun")
 */
func ReadFromSourceTest() {
	//标准输入读取
	//ReadFrom(os.Stdin,10)

	// 从普通文件读取，其中 file 是 os.File 的实例
	//f, _ := os.Open("/data1/www/hu.txt")
	//ReadFrom(f,10)

	// 从字符串读取
	bytes, err := ReadFrom(strings.NewReader("hujianjun"), 9)
	fmt.Printf("读取字节内容【%s】，返回错误信息【%v】", string(bytes), err)
}

/**
 * 读取多个文件内容到[]byte中
 */
func ReadMultiFile(fileNames ...string) []byte {
	data := make([]byte, 0, 128)
	buf := make([]byte, 10)
	readers := make([]io.Reader, 0, len(fileNames))
	for _, fileName := range fileNames {
		file, _ := os.Open(fileName)
		readers = append(readers, file)
	}
	reader := io.MultiReader(readers...)
	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}
	return data
}

/**
 * 写入数据到多个文件中
 */
func WriteMultiFile(content string, fileNames ...string) {
	writers := make([]io.Writer, 0, len(fileNames))
	for _, fileName := range fileNames {
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		writers = append(writers, file)
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte(content))
}
