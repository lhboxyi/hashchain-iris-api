package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

// Md5Sum calculates the md5sum of a stream
// 计算流的md5sum
func Md5Sum(reader io.Reader) (string, error) {
	var returnMD5String string
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// Md5SumString calculates the md5sum of a string
// 计算字符串的md5sum
func Md5SumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return Md5Sum(buffer)
}
