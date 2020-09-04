package util

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	//防止中文乱码的一个库
	"github.com/axgle/mahonia"
)

/**
判断所给路径文件/文件夹是否存在
*/
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
按行读取文件内容
*/
func ReadLineText(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		log.Println(lineText)
	}
}

/**
按文件读取内容
*/
func ReadFileText(filePath string) string {
	byteText, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteText)
}


func ReadAll(filePth string) (content string, err error) {
	file, err := os.Open(filePth)
	if err != nil {
		return "", err
	}
	contentByte, _ := ioutil.ReadAll(file)
	content = string(contentByte)
	defer file.Close()
	return
}

/**
断点续传
 */
func ContinueCopy(srcFile, destDir string) (int, error) {
	// 1. 定义源文件
	fileSrc, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	log.Printf("源文件名称:%s\n", fileSrc.Name())

	// 2. 定义目标文件位置,不存在时自动创建
	destFile := destDir + srcFile[strings.LastIndex(srcFile, "/")+1:]
	fileDest, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	log.Printf("目标文件名称:%s\n", fileDest.Name())

	// 3. 定义零时文件位置,不存在时自动创建(不建议使用ioutil.TempFile(),不便下次找到)
	tempFile := destFile + "_temp"
	fileTemp, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	log.Printf("临时文件名称:%s\n", fileTemp.Name())

	// 4. 关闭文件
	defer fileSrc.Close()
	defer fileDest.Close()
	defer fileTemp.Close()

	// 6. 读取临时文件中的偏移量的值
	tempOffsetStr, err := ioutil.ReadFile(tempFile)
	tempOffSet, err := strconv.ParseInt(string(tempOffsetStr), 10, 64)

	// 7. 本次拷贝的初始位置
	fileSrc.Seek(tempOffSet, io.SeekStart)
	fileDest.Seek(tempOffSet, io.SeekStart)
	data := make([]byte, 1024, 1024)
	countOut := -1           //	读出的总量
	countIn := -1            // 写入的总量
	total := int(tempOffSet) // 总量

	srcRead := bufio.NewReader(fileSrc)
	destWrite := bufio.NewWriter(fileDest)

	// 8. 拷贝文件
	for {
		countOut, err = srcRead.Read(data)
		if err == io.EOF || countOut == 0 {
			log.Printf("文件拷贝完成,总共: %d字节\n", total)
			fileTemp.Close()
			os.Remove(tempFile)
			return 1, nil
		}
		//destFile := bufio.NewWriter(fileDest)
		countIn, err = destWrite.Write(data[:countOut])
		destWrite.Flush()
		total += countIn

		// 9. 将当前复制的偏移量，存储到临时文件
		fileTemp.Seek(0, io.SeekStart)
		fileTemp.WriteString(strconv.Itoa(total))

		// 10. 建设异常情况，突然终止拷贝
		unknownErr(total)
	}
}

// 2M时中断程序
func unknownErr(n int) {
	if n == 1024*2 {
		panic("something has coursed error...")
	}
}

/**
缓冲读取（如果文件比较大的情况下建议是缓冲读取）
*/
func buffRead(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	// 建缓冲区
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		fmt.Println(gbkStr)
	}
}


// 处理乱码
// srcStr：处理的数据
// encoding：数据目前的编码
// dstStr：返回的正常数据
func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	// 创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	// 编码器处理字符串为utf8的字符串
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return
}

/**
判断目录是否存在并删除目录
*/
func RemoveDir(dirPath string) bool {
	b := Exists(dirPath)
	if b {
		removeErr := os.RemoveAll(dirPath)
		if removeErr != nil {
			return false
		}
	}
	return true
}

/**
判断目录下是否包含子目录
*/
func DirWithChildDir(dirPath string) bool {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logrus.Errorf("获取文件失败，错误信息：%s", err.Error())
		return false
	}
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			return true
		}
	}
	return false
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
