package standardApi

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileDemo()  {
	file, err := os.OpenFile("/data1/www/hujianjun/f1.txt", os.O_TRUNC, 0666)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	all, _ := ioutil.ReadAll(file)

	fmt.Println(string(all))
}
